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
	ipv6LinkLocalUni := net.ParseIP("fe80::")
	ipv6Global := net.ParseIP("2000::")
	ipv4LinkLocalUni := net.ParseIP("169.254.0.0")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")

	fmt.Println(ipv6LinkLocalUni.IsLinkLocalUnicast())
	fmt.Println(ipv6Global.IsLinkLocalUnicast())
	fmt.Println(ipv4LinkLocalUni.IsLinkLocalUnicast())
	fmt.Println(ipv4LinkLocalMulti.IsLinkLocalUnicast())

}
