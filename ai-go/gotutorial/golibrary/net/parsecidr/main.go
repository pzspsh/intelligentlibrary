/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:39:49
*/
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ipv4Addr, ipv4Net, err := net.ParseCIDR("192.0.2.1/24")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Addr)
	fmt.Println(ipv4Net)

	ipv6Addr, ipv6Net, err := net.ParseCIDR("2001:db8:a0b:12f0::1/32")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6Addr)
	fmt.Println(ipv6Net)

}
