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
	ipv4DNS := net.ParseIP("8.8.8.8")
	ipv4Lo := net.ParseIP("127.0.0.1")
	ipv6DNS := net.ParseIP("0:0:0:0:0:FFFF:0808:0808")

	fmt.Println(ipv4DNS.Equal(ipv4DNS))
	fmt.Println(ipv4DNS.Equal(ipv4Lo))
	fmt.Println(ipv4DNS.Equal(ipv6DNS))

}
