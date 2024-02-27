/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:51:54
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveIPAddr("ip", "1.1.1.1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Addr", addr.String())
}
