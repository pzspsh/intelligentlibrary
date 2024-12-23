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
	ipv6Lo := net.ParseIP("::1")
	ipv6 := net.ParseIP("ff02::1")
	ipv4Lo := net.ParseIP("127.0.0.0")
	ipv4 := net.ParseIP("128.0.0.0")

	fmt.Println(ipv6Lo.IsLoopback())
	fmt.Println(ipv6.IsLoopback())
	fmt.Println(ipv4Lo.IsLoopback())
	fmt.Println(ipv4.IsLoopback())

}
