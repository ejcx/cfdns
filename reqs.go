package cfdns

import (
	"net/http"
	"time"
)

func newDefaultClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
	}
}

func newReq(reqType string) *DNSRequest {
	return &DNSRequest{
		Client: newDefaultClient(),
		Type:   reqType,
	}
}

// NewA creates a request for an A record.
func NewA() *DNSRequest {
	return newReq("A")
}

// NewAAAA creates a request for an AAAA record.
func NewAAAA() *DNSRequest {
	return newReq("AAAA")
}

// NewCAA creates a request for an CAA record.
func NewCAA() *DNSRequest {
	return newReq("CAA")
}

// NewCDNSKEY creates a request for an CDNSKEY record.
func NewCDNSKEY() *DNSRequest {
	return newReq("CDNSKEY")
}

// NewCDS creates a request for an CDS record.
func NewCDS() *DNSRequest {
	return newReq("CDS")
}

// NewCERT creates a request for an CERT record.
func NewCERT() *DNSRequest {
	return newReq("CERT")
}

// NewCNAME creates a request for an CNAME record.
func NewCNAME() *DNSRequest {
	return newReq("CNAME")
}

// NewDNAME creates a request for an DNAME record.
func NewDNAME() *DNSRequest {
	return newReq("DNAME")
}

// NewIPSECKEY creates a request for an IPSECKEY record.
func NewIPSECKEY() *DNSRequest {
	return newReq("IPSECKEY")
}

// NewLOC creates a request for an LOC record.
func NewLOC() *DNSRequest {
	return newReq("LOC")
}

// NewMX creates a request for an MX record.
func NewMX() *DNSRequest {
	return newReq("MX")
}

// NewNAPTR creates a request for an NAPTR record.
func NewNAPTR() *DNSRequest {
	return newReq("NAPTR")
}

// NewNS a request for an NS record.
func NewNS() *DNSRequest {
	return newReq("NS")
}

// NewPTR a request for an PTR record.
func NewPTR() *DNSRequest {
	return newReq("PTR")
}

// NewRRSIG a request for an RRSIG record.
func NewRRSIG() *DNSRequest {
	return newReq("RRSIG")
}

// NewSOA a request for an SOA record.
func NewSOA() *DNSRequest {
	return newReq("SOA")
}

// NewTLSA a request for an TLSA record.
func NewTLSA() *DNSRequest {
	return newReq("TLSA")
}

// NewTSIG a request for an TSIG record.
func NewTSIG() *DNSRequest {
	return newReq("TSIG")
}

// NewTXT a request for an TXT record.
func NewTXT() *DNSRequest {
	return newReq("TXT")
}

// NewURI a request for an URI record.
func NewURI() *DNSRequest {
	return newReq("URI")
}
