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
	ipv6Multi := net.ParseIP("FF00::")
	ipv6LinkLocalMulti := net.ParseIP("ff02::1")
	ipv6Lo := net.ParseIP("::1")
	ipv4Multi := net.ParseIP("239.0.0.0")
	ipv4LinkLocalMulti := net.ParseIP("224.0.0.0")
	ipv4Lo := net.ParseIP("127.0.0.0")

	fmt.Println(ipv6Multi.IsMulticast())
	fmt.Println(ipv6LinkLocalMulti.IsMulticast())
	fmt.Println(ipv6Lo.IsMulticast())
	fmt.Println(ipv4Multi.IsMulticast())
	fmt.Println(ipv4LinkLocalMulti.IsMulticast())
	fmt.Println(ipv4Lo.IsMulticast())

}
