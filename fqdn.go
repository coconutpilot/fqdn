// Package fqdn - Best effort to return the machine's FQDN.
//
// The current method is to lookup the machine's IP in DNS.  This
// requires your network has working forward and reverse DNS.
//
// For multi-homed machines this library returns the FQDN for the
// default local IP.  The default local IP is the source IP used when creating
// a dummy connection to 8.8.8.8.
package fqdn

import (
	"fmt"
	"net"
	"strings"
)

// support mocking in tests
var defaultaddr = "8.8.8.8:8"
var getDefaultIP = _getDefaultIP
var getDNSFQDN = _getDNSFQDN
var lookupAddr = net.LookupAddr
var lookupHost = net.LookupHost

func _getDefaultIP() (net.IP, error) {
	conn, err := net.Dial("udp", defaultaddr)
	if err != nil {
		return nil, fmt.Errorf("Failed looking up my IP address: %v", err)
	}
	defer conn.Close()
	return (conn.LocalAddr().(*net.UDPAddr)).IP, nil
}

func verifyFQDN(fqdn string) (bool, error) {
	ips, err := lookupHost(fqdn)
	if err != nil {
		return false, err
	}
	for _, ip := range ips {
		hosts, err := lookupAddr(ip)
		if err != nil {
			continue
		}
		for _, host := range hosts {
			if fqdn == host {
				return true, nil
			}
		}
	}
	return false, nil
}

func _getDNSFQDN() (string, error) {
	ip, err := getDefaultIP()
	if err != nil {
		return "", err
	}

	hosts, err := lookupAddr(ip.String())
	if err != nil {
		return "", fmt.Errorf("Error looking up FQDN: %v", err)
	}
	if len(hosts) == 0 {
		return "", fmt.Errorf("Unable to lookup FQDN: %v", err)
	}
	return hosts[0], nil
}

var _fqdn = ""

// FQDN - Returns FQDN or error if unable to determine
func FQDN() (string, error) {
	if _fqdn != "" {
		return _fqdn, nil
	}

	fqdn, err := getDNSFQDN()
	if err != nil {
		return "", fmt.Errorf("Error resolving rDNS: %v", err)
	}

	ok, err := verifyFQDN(fqdn)
	if err != nil {
		return "", fmt.Errorf("Error resolving DNS: %v", err)
	}
	if !ok {
		return "", fmt.Errorf("FQDN failed verification: %v", err)
	}

	_fqdn = strings.TrimSuffix(fqdn, ".")
	return _fqdn, nil
}
