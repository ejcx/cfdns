package cfdns

import (
	"testing"
)

func TestARequest(t *testing.T) {
	dc := NewA()
	resp, err := dc.Do("cloudflare.com")
	if err != nil {
		t.Fatalf("Could not make dns request: %s", err)
	}
	if resp.Status != 0 {
		t.Fatalf("Cloudflare having major outage: %d", resp.Status)
	}
}

func TestAAAARequest(t *testing.T) {
	dc := NewAAAA()
	resp, err := dc.Do("cloudflare.com")
	if err != nil {
		t.Fatalf("Could not make AAAA dns request: %s", err)
	}
	if resp.Status != 0 {
		t.Fatalf("Cloudflare having major outage: %d", resp.Status)
	}
}
