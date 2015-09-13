package main

import (
	"github.com/miekg/dns"
	"fmt"
	"os"
	"net"
        "strings"
)

var dnsTypeValueByName = map[string]uint16 {
	"a": dns.TypeA,
	"any": dns.TypeANY,
	"cname": dns.TypeCNAME,
	"hinfo": dns.TypeHINFO,
	"mx": dns.TypeMX,
	"ns": dns.TypeNS,
	"ptr": dns.TypePTR,
	"soa": dns.TypeSOA,
	"txt": dns.TypeTXT,
}

func printError(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "dnsq: ERROR: ")
	fmt.Fprintf(os.Stderr, format, a...)
}

func printUsage() {
	const usage = "dnsq: Usage: dnsq TYPE NAME SERVER\n"
	os.Stderr.Write([]byte(usage))
}

func main() {
	ret := 0
	defer func() { os.Exit(ret) }()

	if len(os.Args) != 4 {
		ret = 100
		printUsage()
		return
	}
	q_type, ok := dnsTypeValueByName[strings.ToLower(os.Args[1])]
	if !ok {
		ret = 100
		printError("Unknown type: %v\n", os.Args[1])
		return
	}
	q_name := os.Args[2]
	ns := os.Args[3]

	c := new(dns.Client)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(q_name), q_type)
	m.RecursionDesired = false

	r, _, err := c.Exchange(m, net.JoinHostPort(ns, "53"))
	if r == nil {
		printError("FIXME: %s\n", err.Error())
	}

	if r.Rcode != dns.RcodeSuccess {
		printError("FIXME: Print error detailss\n")
	}
	// Stuff must be in the answer section
	for _, a := range r.Answer {
		fmt.Printf("%v\n", a)
	}
}

