package fqdn

import (
	"errors"
	"net"
	"testing"
)

func TestGetDefaultIP(t *testing.T) {
	_fqdn = ""
	ip, err := getDefaultIP()
	if err != nil {
		t.Error(err)
	}
	t.Logf("My IP: %v", ip)
}

func TestFQDN(t *testing.T) {
	_fqdn = ""
	fqdn, err := FQDN()
	if err != nil {
		t.Error(err)
	}
	t.Logf("My FQDN: %v", fqdn)
}

func TestCachedFQDN(t *testing.T) {
	_fqdn = "bar.example.com"
	fqdn, err := FQDN()
	if err != nil {
		t.Error(err)
	}
	if fqdn != _fqdn {
		t.Errorf("Expected: %v, got: %v", _fqdn, fqdn)
	}
	t.Logf("My FQDN: %v", fqdn)
}

func TestFQDNLocalhost(t *testing.T) {
	_fqdn = ""
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
	_fqdn = ""
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
	_fqdn = ""
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
	_fqdn = ""
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
	_fqdn = ""
	defaultaddr = "8.8.8.8:8"
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
	_fqdn = ""
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

func TestFQDNError5(t *testing.T) {
	_fqdn = ""
	lookupAddr = net.LookupAddr
	getDefaultIP = _getDefaultIP
	lookupHost = func(addr string) (names []string, err error) {
		return []string{}, errors.New("Testing net.LookupHost() failure")
	}
	fqdn, err := FQDN()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: ''", fqdn)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}
}

func TestFQDNError6(t *testing.T) {
	_fqdn = ""
	lookupHost = func(addr string) (names []string, err error) {
		return []string{"1.2.3.4", ""}, nil
	}
	lookupAddr = func(addr string) (names []string, err error) {
		if addr == "" {
			return []string{}, errors.New("Testing rDNS failure")
		}
		return []string{"bar.example.com"}, nil
	}
	getDNSFQDN = func() (string, error) {
		return "foo.example.com", nil
	}

	fqdn, err := FQDN()
	if err == nil {
		t.Errorf("Expected an error here, got: %v, expected: ''", fqdn)
	} else {
		t.Logf("Testing error handling, received: %v", err)
	}
}
