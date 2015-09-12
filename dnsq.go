package main

import (
	"github.com/miekg/dns"
	"fmt"
	"os"
	"net"
        "log"
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
}

func printUsage() {
	const usage = "dnsq: Usage: dnsq TYPE NAME SERVER\n"
	os.Stderr.Write([]byte(usage))
	os.Exit(100)
}

func main() {
	ret := 0
	defer func() { os.Exit(ret) }()

	if len(os.Args) != 4 {
		printUsage()
	}
	q_type, _ := dnsTypeValueByName[os.Args[1]]
	q_name := os.Args[2]
	ns := os.Args[3]

	c := new(dns.Client)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(q_name), q_type)
	m.RecursionDesired = false

	r, _, err := c.Exchange(m, net.JoinHostPort(ns, "53"))
	if r == nil {
		log.Fatalf("*** error: %s\n", err.Error())
	}

	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("ERROR: FIXME: Print error detailss\n")
	}
	// Stuff must be in the answer section
	for _, a := range r.Answer {
		fmt.Printf("%v\n", a)
	}
}

