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
	ipv6Private := net.ParseIP("fc00::")
	ipv6Public := net.ParseIP("fe00::")
	ipv4Private := net.ParseIP("10.255.0.0")
	ipv4Public := net.ParseIP("11.0.0.0")

	fmt.Println(ipv6Private.IsPrivate())
	fmt.Println(ipv6Public.IsPrivate())
	fmt.Println(ipv4Private.IsPrivate())
	fmt.Println(ipv4Public.IsPrivate())

}
