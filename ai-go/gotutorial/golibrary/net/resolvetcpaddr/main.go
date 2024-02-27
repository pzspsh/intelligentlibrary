/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:51:22
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "1.1.1.1:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("AddrPort", addr.AddrPort())
	fmt.Println("network", addr.Network())
	fmt.Println("Addr", addr.String())
}
