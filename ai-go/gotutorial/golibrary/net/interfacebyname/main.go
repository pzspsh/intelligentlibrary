/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:50:20
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	en0, err := net.InterfaceByName("en0")
	if err != nil {
		fmt.Println(err)
		// error handling
	}
	en1, err := net.InterfaceByIndex(911)
	if err != nil {
		fmt.Println(err)
		// error handling
	}
	group := net.IPv4(224, 0, 0, 250)
	fmt.Println(en0, en1, group)
}
