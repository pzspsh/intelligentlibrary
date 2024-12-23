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
	// This mask corresponds to a /31 subnet for IPv4.
	fmt.Println(net.CIDRMask(31, 32))

	// This mask corresponds to a /64 subnet for IPv6.
	fmt.Println(net.CIDRMask(64, 128))

}
