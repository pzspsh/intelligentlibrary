/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 15:39:49
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6LinkLocalMulti := net.ParseIP("ff02::2")
	ipv6LinkLocalUni := net.ParseIP("fe80::")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	ipv4LinkLocalUni := net.ParseIP("169.254.0.0")

	fmt.Println(ipv6LinkLocalMulti.IsLinkLocalMulticast())
	fmt.Println(ipv6LinkLocalUni.IsLinkLocalMulticast())
	fmt.Println(ipv4LinkLocalMulti.IsLinkLocalMulticast())
	fmt.Println(ipv4LinkLocalUni.IsLinkLocalMulticast())

}
