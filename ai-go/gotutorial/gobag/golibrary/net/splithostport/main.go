/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:20:29
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	host, port, _ := net.SplitHostPort("192.168.110.66:8080")
	fmt.Println("host:", host, "port:", port)
	fmt.Println("拼接", net.JoinHostPort(host, port))
}
