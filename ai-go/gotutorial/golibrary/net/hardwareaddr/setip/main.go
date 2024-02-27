/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:17:17
*/
package main

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func main() {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Local IP: ", ipnet.IP.String())
			}
		}
	}

	link, err := netlink.LinkByName("eth0")
	if err != nil {
		panic(err)
	}

	addr, err := netlink.ParseAddr("192.168.0.2/24")
	if err != nil {
		panic(err)
	}

	if err := netlink.AddrAdd(link, addr); err != nil {
		panic(err)
	}

	addrs, err = net.InterfaceAddrs()
	if err != nil {
		fmt.Println("Error:", err)
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("Local IP: ", ipnet.IP.String())
			}
		}
	}
}
