package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	interfaces, err := net.Interfaces()
	for _, inter := range interfaces {
		fmt.Println(inter.HardwareAddr.String())
		if err != nil {
			continue
		}
		return
	}
	return
}

// CheckNetIface 检查网卡
func CheckNetIface(ethName string) bool {
	return strings.HasPrefix(ethName, "en") || strings.HasPrefix(ethName, "eth") || strings.HasPrefix(ethName, "wl")
}

func netMacs() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ifaces []string
	for _, inter := range interfaces {
		if CheckNetIface(inter.Name) {
			ifaces = append(ifaces, inter.HardwareAddr.String())
		}
	}

	return ifaces, nil
}
