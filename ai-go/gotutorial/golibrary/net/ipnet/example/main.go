/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:08:18
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	mask := net.IPv4Mask(byte(255), byte(255), byte(255), byte(0))
	ip := net.ParseIP("192.168.1.125").Mask(mask)
	in := &net.IPNet{IP: ip, Mask: mask}
	fmt.Println(in) //  192.168.1.0/24
}
