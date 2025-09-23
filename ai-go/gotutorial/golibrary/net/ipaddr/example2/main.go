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
	fmt.Println("network", addr.Network())
	fmt.Println("Addr", addr.String())
}
