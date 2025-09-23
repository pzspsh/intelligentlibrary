package main

import (
	"fmt"
	"net"
)

func main() {
	ipv6 := net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ipv4 := net.IPv4(10, 255, 0, 0)

	fmt.Println(ipv6.To4())
	fmt.Println(ipv4.To4())

}
