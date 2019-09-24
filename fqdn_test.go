package fqdn

import (
	"net"
	"testing"
)

func TestGetDefaultIP(t *testing.T) {
	ip, err := getDefaultIP()
	if err != nil {
		t.Error(err)
	}
	t.Logf("My IP: %v", ip)
}

func TestFQDN(t *testing.T) {
	fqdn, err := FQDN()
	if err != nil {
		t.Error(err)
	}
	t.Logf("My FQDN: %v", fqdn)
}

func TestFQDNLocalhost(t *testing.T) {
	defaultaddr = "127.0.0.1:1"
	fqdn, err := FQDN()
	if err != nil {
		t.Error(err)
	}
	if fqdn != "localhost" {
		t.Errorf("Unexpected return value, got: %v, expected: localhost", fqdn)
	}
}

// this breaks net.Dial()
func TestGetDefaultIPError(t *testing.T) {
	defaultaddr = ""
	ip, err := getDefaultIP()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: nil", ip)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}
}

// this breaks net.Dial()
func TestFQDNError(t *testing.T) {
	defaultaddr = ""
	fqdn, err := FQDN()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: ''", fqdn)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}

}

// this breaks net.LookupAddr
func TestFQDNError2(t *testing.T) {
	getDefaultIP = func() (net.IP, error) {
		var ip = net.IP{}
		return ip, nil
	}
	fqdn, err := FQDN()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: ''", fqdn)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}

}

// this tests an IP we can't do PTR DNS on
func TestFQDNError3(t *testing.T) {
	getDefaultIP = func() (net.IP, error) {
		var ip = net.IP{0, 0, 0, 0}
		return ip, nil
	}
	fqdn, err := FQDN()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: ''", fqdn)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}
}

// this simulates a broken net.LookupAddr by returning an empty slice
func TestFQDNError4(t *testing.T) {
	lookupAddr = func(addr string) (names []string, err error) {
		return []string{}, nil
	}
	fqdn, err := FQDN()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: ''", fqdn)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}
}
