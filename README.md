# FQDN - Get the machine's FQDN

[![Build Status](https://travis-ci.org/coconutpilot/)](https://travis-ci.org/coconutpilot/fqdn)

### What is a FQDN?

A FQDN is a DNS name that resolves to one of the host's IPs.

A FQDN may be different from the name that the machine's kernel returns
from the hostname() syscall.

A multihomed machine may have multiple FQDNs, one per IP.  This library
returns the FQDN for the default local IP.


## SYNOPSIS

    myfqdn, err := fqdn.FQDN()
    if err != nil {
        // DNS broken
    }

## DOCUMENTATION

See fqdn.go for usage.


## LICENSE

Copyright 2019 David Sparks.  See the file LICENSE.txt included with the 
FQDN distribution for details.
