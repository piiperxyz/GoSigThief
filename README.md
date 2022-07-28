# GoSigThief

`Golang` version of [sigthief](https://github.com/secretsquirrel/SigThief)

made for [AniYa](https://github.com/piiperxyz/AniYa)

```
sigthief.exe -h
Usage of sigthief.exe:
  -a    add signature to file
  -c string
        certfile
  -i string
        inputfile
  -o string
        outputfile
  -s    save cert to disk
  -show
        show example
```

```
sigthief.exe -show
Save cert: sigthief.exe -i MSbuild.exe -s -o MSbuild.cert
Add cert from certfile: sigthief.exe -a -i evil.exe -o evil-sign.exe -c MSbuild.cert
You can also use exe to add cert: sigthief.exe -a -i evil.exe -o evil-sign.exe -c MSbuild.exe
```

