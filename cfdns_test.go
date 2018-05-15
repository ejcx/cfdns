package cfdns

import (
	"fmt"
	"log"
	"testing"
)

// This is an example of how to make an A request.
func ExampleExamples_ARequest() {
	resp, err := NewA().Do("cloudflare.com")
	if err != nil {
		log.Printf("Could not make DNS Request: %s", err)
		return
	}
	if !resp.OK() {
		log.Printf("Unexpected DNS Response: %s", err)
		return
	}
	fmt.Println(resp.Answer[0].Data)
}

// This is an example of how to make a TXT request.
func ExampleExamples_TXTRequest() {
	resp, err := NewTXT().Do("cloudflare.com")
	if err != nil {
		log.Printf("Could not make DNS Request: %s", err)
		return
	}
	if !resp.OK() {
		log.Printf("Unexpected DNS Response: %s", err)
		return
	}
	fmt.Println(resp.Answer[0].Data)
}

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
