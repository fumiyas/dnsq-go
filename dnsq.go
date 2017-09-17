// djbdns's dnsq(1) clone implemented in Golang
//
// Copyright (C) 2015-2017 SATOH Fumiyasu @ OSS Technology Corp., Japan
// License: Go
// Development home: <https://github.com/fumiyas/dnsq-go>
// Author's home: <https://fumiyas.github.io/>

package main

import (
	"fmt"
	"github.com/miekg/dns"
	"net"
	"os"
	"strings"
)

var dnsTypeValueByName = map[string]uint16{
	"a":     dns.TypeA,
	"aaaa":  dns.TypeAAAA,
	"any":   dns.TypeANY,
	"axfe":  dns.TypeAXFR,
	"cname": dns.TypeCNAME,
	"hinfo": dns.TypeHINFO,
	"key":   dns.TypeKEY,
	"mx":    dns.TypeMX,
	"ns":    dns.TypeNS,
	"ptr":   dns.TypePTR,
	"rp":    dns.TypeRP,
	"sig":   dns.TypeSIG,
	"soa":   dns.TypeSOA,
	"txt":   dns.TypeTXT,
	"srv":   dns.TypeSRV,
}

func printError(format string, a ...interface{}) {
	fmt.Fprint(os.Stderr, "dnsq: ERROR: ")
	fmt.Fprintf(os.Stderr, format, a...)
}

func printUsage() {
	const usage = "dnsq: Usage: dnsq TYPE NAME SERVER\n"
	os.Stderr.Write([]byte(usage))
}

func dnsRRToString(rr dns.RR) string {
	parts := strings.SplitN(rr.String(), "\t", 5)
	return parts[0] + " " + parts[1] + " " + parts[3] + " " + parts[4]
}

func main() {
	ret := 0
	defer func() { os.Exit(ret) }()

	if len(os.Args) != 4 {
		ret = 100
		printUsage()
		return
	}

	qType, ok := dnsTypeValueByName[strings.ToLower(os.Args[1])]
	if !ok {
		ret = 100
		printError("Unknown type: %v\n", os.Args[1])
		return
	}
	qName := os.Args[2]
	ns := os.Args[3]

	c := new(dns.Client)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(qName), qType)
	m.RecursionDesired = false

	fmt.Printf("%v %s:\n", qType, qName)

	r, _, err := c.Exchange(m, net.JoinHostPort(ns, "53"))
	if r == nil {
		printError("DNS client: %s\n", err.Error())
	}

	// FIXME: How to get a response size?
	fmt.Printf("%d+%d+%d+%d records",
		len(r.Question),
		len(r.Answer),
		len(r.Ns),
		len(r.Extra),
	)
	if r.Response {
		fmt.Print(", response")
	}
	if r.Opcode == 0xF {
		fmt.Print(", weird op")
	}
	if r.Authoritative {
		fmt.Print(", authoritative")
	}
	if r.Truncated {
		fmt.Print(", truncated")
	}
	if r.RecursionDesired {
		fmt.Print(", weird rd")
	}
	if r.RecursionAvailable {
		fmt.Print(", weird ra")
	}
	switch r.Rcode {
	case dns.RcodeSuccess:
		fmt.Print(", noerror")
	case dns.RcodeNameError:
		fmt.Print(", nxdomain")
	case dns.RcodeNotImplemented:
		fmt.Print(", notimp")
	case dns.RcodeRefused:
		fmt.Print(", refused")
	default:
		fmt.Print(", weird rcode")
	}
	if r.Zero {
		fmt.Print(", weird z")
	}
	fmt.Print("\n")

	for _, rr := range r.Question {
		fmt.Printf("query: %v %s\n", rr.Qtype, rr.Name)
	}
	for _, rr := range r.Answer {
		fmt.Printf("answer: %v\n", dnsRRToString(rr))
	}
	for _, rr := range r.Ns {
		fmt.Printf("authority: %v\n", dnsRRToString(rr))
	}
	for _, rr := range r.Extra {
		fmt.Printf("additional: %v\n", dnsRRToString(rr))
	}
}
