// Package cfdns is a DNS client that uses the Cloudflare 1.1.1.1 public resolver.
// It uses DNS over HTTPS instead of using the DNS protocol.
package cfdns

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	cfdns = "cloudflare-dns.com"

	// NoError [RFC1035]
	NoError = 0
	// FormErr [RFC1035]
	FormErr = 1
	// ServFail [RFC1035]
	ServFail = 2
	// NXDomain [RFC1035]
	NXDomain = 3
	// NotImp [RFC1035]
	NotImp = 4
	// Refused [RFC1035]
	Refused = 5
	// YXDomain [RFC2136][RFC6672]
	YXDomain = 6
	// YXRRSet [RFC2136]
	YXRRSet = 7
	// NXRRSet [RFC2136]
	NXRRSet = 8
	// NotAuth [RFC2136] [RFC2845]
	NotAuth = 9
	// NotZone [RFC2136]
	NotZone = 10
	// BADVERS [RFC6891]
	BADVERS = 16
	// BADSIG [RFC2845]
	BADSIG = 16
	// BADKEY [RFC2845]
	BADKEY = 17
	// BADTIME [RFC2845]
	BADTIME = 18
	// BADMODE [RFC2930]
	BADMODE = 19
	// BADNAME [RFC2930]
	BADNAME = 20
	// BADALG [RFC2930]
	BADALG = 21
	// BADTRUNC [RFC4635]
	BADTRUNC = 22
	// BADCOOKIE [RFC7873]
	BADCOOKIE = 23
)

// DNSResponse is a JSON form of a DNS response.
// It contains the raw fields as they are abbreviated
// in the Cloudflare DNS over HTTPS response.
type DNSResponse struct {
	Status    int         `json:"Status"`
	TC        bool        `json:"TC"`
	RD        bool        `json:"RD"`
	RA        bool        `json:"RA"`
	AD        bool        `json:"AD"`
	CD        bool        `json:"CD"`
	Question  []Question  `json:"Question"`
	Authority []Authority `json:"Authority"`
	Answer    []Answer    `json:"Answer"`
}

// Question is the question part of the DNS request echo'd back in the response.
type Question struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

// Answer is part of the DNS response. It contains the meat of the response.
type Answer struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

// Authority is part of the DNS response.
type Authority struct {
	Name string `json:"name"`
	Type int    `json:"type"`
	TTL  int    `json:"TTL"`
	Data string `json:"data"`
}

// DNSRequest is all the information needed to make a DNS Request.
// It is possible to set your own HTTP Client if you have special
// requirements, or set the DNSRequest to a type that is not
// supported by this library.
type DNSRequest struct {
	Client *http.Client
	Type   string
}

// Do will make a DNS Request to the host of your choice.
// *Note:* Do not pass a URL to Do. Only a hostname
func (d *DNSRequest) Do(host string) (*DNSResponse, error) {
	u := fmt.Sprintf("https://%s/dns-query?ct=application/dns-json&name=%s&type=%s", cfdns, host, d.Type)
	// If a scheme is passed then we can catch it here because
	// the URL `https://https://` will not parse properly.
	_, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	dr := new(DNSResponse)
	err = json.Unmarshal(buf, dr)
	return dr, err
}

// OK returns true if the DNS Response Status is 0 indicated there
// was No Error.
func (d *DNSResponse) OK() bool {
	return d.Status == NoError
}

// Truncated is true if true the truncated bit was set. This happens
// when the DNS answer is larger than a single UDP or TCP packet. TC
// will almost always be false with Cloudflare DNS over HTTPS because
// Cloudflare supports the maximum response size.
func (d *DNSResponse) Truncated() bool {
	return d.TC
}

// RecursiveDesired is true if the Recursion Desired bit was set.
// This should always be set to true for Cloudflare DNS over HTTPS.
func (d *DNSResponse) RecursiveDesired() bool {
	return d.RD
}

// RecursiveAvailable is true if the Recursion Available bit was set.
// This should always be set to true for Cloudflare DNS over HTTPS.
func (d *DNSResponse) RecursiveAvailable() bool {
	return d.RA
}

// AuthenticatedWithDNSSEC is true if every record in the answer
// was verified with DNSSEC.
func (d *DNSResponse) AuthenticatedWithDNSSEC() bool {
	return d.AD
}

// ClientDisabledDNSSEC is true if the client asked to disable DNSSEC
// validation. In this case, Cloudflare will still fetch the DNSSEC
// related records, but it will not attempt to validate the records.
func (d *DNSResponse) ClientDisabledDNSSEC() bool {
	return d.CD
}
