djbdns's dnsq(1) and dnsqr(1) clone implemented in Golang
======================================================================

  * Copyright (C) 2015 SATOH Fumiyasu @ OSS Technology Corp., Japan
  * Development home: <https://github.com/fumiyas/dnsq-go>
  * Author's home: <https://fumiyas.github.io/>

What's this?
----------------------------------------------------------------------

だいたい完成。

How to build native binaries:

```console
$ go build dnsq.go
$ go build dnsqr.go
```

How to build Windows binaries on non-Windows environment (cross build):

```console
$ GOOS=windows go build dnsq.go
$ GOOS=windows go build dnsqr.go
```

