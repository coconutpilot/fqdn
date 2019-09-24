// Package fqdn - Best effort to return the machine's FQDN.  This requires
// your network has working forward and reverse DNS.
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
var lookupAddr = net.LookupAddr

func _getDefaultIP() (net.IP, error) {
	conn, err := net.Dial("udp", defaultaddr)
	if err != nil {
		return nil, fmt.Errorf("Failed looking up my IP address: %v", err)
	}
	defer conn.Close()
	return (conn.LocalAddr().(*net.UDPAddr)).IP, nil
}

// FQDN - Returns FQDN or error if unable to determine
func FQDN() (string, error) {
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
	return strings.TrimSuffix(hosts[0], "."), nil
}
