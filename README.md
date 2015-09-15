djbdns's dnsq(1) and dnsqr(1) clone implemented in Golang
======================================================================

  * Copyright (C) 2015 SATOH Fumiyasu @ OSS Technology Corp., Japan
  * Development home: <https://github.com/fumiyas/dnsq-go>
  * Author's home: <https://fumiyas.github.io/>

What's this?
----------------------------------------------------------------------

djbdns's dnsq(1) and dnsqr(1) clone implemented in Golang

  * dnsq
    * Sends a non-recursive DNS query to DNS contents server
  * dnsqr
    * Sends a recursive DNS query to DNS cache server
  * djbdns
    * http://cr.yp.to/djbdns.html

How to build
----------------------------------------------------------------------

How to build native binaries:

```console
$ go build dnsq.go
$ go build dnsqr.go
```

How to build Windows binaries on non-Windows environment (cross build):

```console
$ GOOS=windows GOARCH=386 go build dnsq.go
$ GOOS=windows GOARCH=386 go build dnsqr.go
```

How to use
----------------------------------------------------------------------

```console
$ dnsq a www.google.com a.root-servers.net
$ dnsqr a www.osstech.co.jp
$ dnsqr a www.xvideos.com 8.8.8.8
$ dnsqr a www.xvideos.com your-full-service-resolver.example.jp
```

