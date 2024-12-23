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
	fmt.Println(net.IPv4Mask(255, 255, 255, 0))
}
