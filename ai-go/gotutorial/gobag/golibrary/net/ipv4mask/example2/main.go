/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 17:43:33
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	i, err := net.InterfaceByName("en0") // en0 表示当前网络接口
	if err != nil {
		fmt.Println(err)
		return
	}
	addrs, err := i.Addrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("IP地址:", ipnet.IP.String())
			}
		}
	}
	addrs2, _ := i.Addrs()
	for _, addr := range addrs2 {
		fmt.Println("地址:", addr)
	}
	fmt.Println("MAC地址:", i.HardwareAddr)
	fmt.Println("子网掩码:", net.IPv4Mask(255, 255, 255, 0).String())
	addr, _ := i.Addrs()
	fmt.Println("网关地址:", addr[0].(*net.IPNet).IP)
}
