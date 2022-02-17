//go:build windows
// +build windows

package gateway

import (
	"net"
	"os/exec"
	"strings"
	"syscall"
)

func discoverGatewayOSSpecific() (ip net.IP, err error) {
	routeCmd := exec.Command("route", "print", "0.0.0.0")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseWindowsGatewayIP(output)
}

func discoverGatewayInterfaceOSSpecific() (ip net.IP, err error) {
	routeCmd := exec.Command("route", "print", "0.0.0.0")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return parseWindowsInterfaceIP(output)
}

func discoverGatewayInterfaceNameOSSpecific() (name string, err error) {
	routeCmd := exec.Command("route", "print", "0.0.0.0")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	devs, err := getWindowsDeviceMap()
	if err != nil {
		return "", err
	}
	return parseWindowsInterfaceNameIP(output, devs)
}

func getWindowsDeviceMap() (devs map[string]string, err error) {
	devs = make(map[string]string)
	routeCmd := exec.Command("getmac")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	lines := strings.Split(string(output), "\n")
	for _, line := range lines[3:] {
		items := strings.Fields(line)
		if len(items) > 1 {
			devs[items[0]] = items[1]
		}
	}
	return
}
