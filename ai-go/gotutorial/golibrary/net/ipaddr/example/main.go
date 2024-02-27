/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:52:41
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
	fmt.Println("Addr", addr.String())
}
