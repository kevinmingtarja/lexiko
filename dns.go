package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/miekg/dns"
	"net"
)

type handler struct {
	db *redis.Client
}

func (h *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)

	switch r.Question[0].Qtype {
	case dns.TypeA:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := h.getAddress(domain)
		if ok {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{ Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60 },
				A: net.ParseIP(address),
			})
		}
	}

	err := w.WriteMsg(&msg)
	if err != nil {
		return
	}
}