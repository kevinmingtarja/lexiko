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
	case dns.TypeCNAME:
		msg.Authoritative = true
		domain := msg.Question[0].Name
		address, ok := h.getAddress(domain)
		if ok {
			msg.Answer = append(msg.Answer, &dns.CNAME{
				Hdr:    dns.RR_Header{ Name: domain, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 60 },
				Target: address,
			})
		}
	}

	err := w.WriteMsg(&msg)
	if err != nil {
		return
	}
}

func prepareAnswerMsg(req *dns.Msg, answers []dns.RR) *dns.Msg {
	answerMsg := new(dns.Msg)
	answerMsg.Id = req.Id
	answerMsg.Response = true
	answerMsg.Authoritative = true
	answerMsg.Question = req.Question
	answerMsg.Answer = answers
	answerMsg.Rcode = dns.RcodeSuccess
	answerMsg.Extra = []dns.RR{}
	return answerMsg
}

func prepareFailureMsg(req *dns.Msg) *dns.Msg {
	failMsg := new(dns.Msg)
	failMsg.Id = req.Id
	failMsg.Response = true
	failMsg.Authoritative = true
	failMsg.Question = req.Question
	failMsg.Rcode = dns.RcodeNameError
	return failMsg
}

//func answerA(q *dns.Question, v *DNSValue) dns.RR {
//	answer := new(dns.A)
//	answer.Header().Name = q.Name
//	answer.Header().Rrtype = dns.TypeA
//	answer.Header().Class = dns.ClassINET
//	answer.A = net.ParseIP(v.Value)
//	return answer
//}
//
//func answerCNAME(q *dns.Question, v *DNSValue) (dns.RR, string) {
//	answer := new(dns.CNAME)
//	answer.Header().Name = q.Name
//	answer.Header().Rrtype = dns.TypeCNAME
//	answer.Header().Class = dns.ClassINET
//	answer.Target = strings.TrimSuffix(v.Value, ".") + "."
//	return answer, answer.Target
//}