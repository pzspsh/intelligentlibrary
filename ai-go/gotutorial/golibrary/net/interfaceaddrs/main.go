/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:21:45
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.String())
		fmt.Println(addr.Network())
	}
}
