/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:10:15
*/
package main

import (
	"fmt"
	"net"
)

func getMacAddress(iface net.Interface) net.HardwareAddr {
	addr, err := iface.Addrs()
	if err != nil {
		fmt.Println("get addrs err:", err)
		return nil
	}
	for _, addr := range addr {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLinkLocalUnicast() {
			return iface.HardwareAddr
		}
	}
	return nil
}

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("get interfaces err:", err)
		return
	}
	for _, iface := range ifaces {
		fmt.Printf("Interface:%v\n", iface.Name)
		fmt.Printf("\tMAC Address:%v\n", getMacAddress(iface))
	}
}
