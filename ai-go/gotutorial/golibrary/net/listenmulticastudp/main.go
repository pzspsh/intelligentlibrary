/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:56:26
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.ListenMulticastUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP("224.0.0.1"), Port: 1234})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(l)
}
