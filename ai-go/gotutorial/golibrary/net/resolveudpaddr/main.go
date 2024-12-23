/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:50:53
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("tcp", "1.1.1.1:53")
	if err != nil {
		panic(err)
	}
	fmt.Println("Addr", addr.String())
}
