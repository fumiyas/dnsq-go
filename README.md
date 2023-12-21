djbdns's dnsq(1) and dnsqr(1) clone implemented in Golang
======================================================================

  * Copyright (C) 2015-2023 SATOH Fumiyasu @ OSS Technology Corp., Japan
  * License: Go
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

Download
---------------------------------------------------------------------

Binary files are here for Windows:

  * https://github.com/fumiyas/dnsq-go/releases

How to install
----------------------------------------------------------------------

```console
$ go install github.com/fumiyas/dnsq-go/cmd/dnsq@latest
$ go install github.com/fumiyas/dnsq-go/cmd/dnsqr@latest
$ ls ~/go/bin/dnsq*
...
```

How to build
----------------------------------------------------------------------

How to build native binaries:

```console
$ git clone https://github.com/fumiyas/dnsq-go.git
$ cd dnsq-go
$ make
...
$ ls build/bin
...
```

How to build Windows binaries on non-Windows environment (cross build):

```console
$ GOOS=windows go build ./cmd/dnsq
$ GOOS=windows go build ./cmd/dnsqr
$ ls *.exe
...
```

or:

```console
$ GOOS=windows GOARCH=386 go build ./cmd/dnsq
$ GOOS=windows GOARCH=386 go build ./cmd/dnsqr
$ ls *.exe
...
```

How to use
----------------------------------------------------------------------

```console
$ dnsq a www.google.com a.root-servers.net
$ dnsqr a www.osstech.co.jp
$ dnsqr a www.xvideos.com 8.8.8.8
$ dnsqr a www.xvideos.com your-full-service-resolver.example.jp
```

TODO
----------------------------------------------------------------------

  * Add an option to enable/disable TCP, UDP, EDNS0 and so on.
  * Add an option to specify EDNS0 buffer size.
  * Support DNSSEC.
