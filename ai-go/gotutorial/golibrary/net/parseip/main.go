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
	fmt.Println(net.ParseIP("192.0.2.1"))
	fmt.Println(net.ParseIP("2001:db8::68"))
	fmt.Println(net.ParseIP("192.0.2"))

}
