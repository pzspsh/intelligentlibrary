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
	ip := net.ParseIP("192.0.2.1")
	fmt.Println(ip.DefaultMask())

}
