/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:42:26
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	mask := net.IPv4Mask(byte(255), byte(255), byte(255), byte(0))
	ip := net.ParseIP("192.168.1.125").Mask(mask)
	in := &net.IPNet{ip, mask}
	fmt.Println(in) //  192.168.1.0/24
}
