//go:build !darwin && !linux && !windows && !solaris && !freebsd
// +build !darwin,!linux,!windows,!solaris,!freebsd

package gateway

import (
	"net"
)

func discoverGatewayOSSpecific() (ip net.IP, err error) {
	return ip, errNotImplemented
}

func discoverGatewayInterfaceOSSpecific() (ip net.IP, err error) {
	return nil, errNotImplemented
}

func discoverGatewayInterfaceNameOSSpecific() (name string, err error) {
	return "", errNotImplemented
}
