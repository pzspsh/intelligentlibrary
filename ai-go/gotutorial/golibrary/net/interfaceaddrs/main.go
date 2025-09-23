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
