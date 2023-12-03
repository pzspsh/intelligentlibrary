/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 16:51:33
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	ipStr := "192.168.0.1"
	ip := net.ParseIP(ipStr)
	if ip == nil {
		fmt.Println("无效的 IP 地址")
		return
	}
	ipAddr := &net.IPNet{
		IP:   ip,
		Mask: ip.DefaultMask(),
	}
	fmt.Println(ipAddr)
}
