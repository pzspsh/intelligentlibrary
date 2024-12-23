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
	ipv6InterfaceLocalMulti := net.ParseIP("ff01::1")
	ipv6Global := net.ParseIP("2000::")
	ipv4 := net.ParseIP("255.0.0.0")

	fmt.Println(ipv6InterfaceLocalMulti.IsInterfaceLocalMulticast())
	fmt.Println(ipv6Global.IsInterfaceLocalMulticast())
	fmt.Println(ipv4.IsInterfaceLocalMulticast())

}
