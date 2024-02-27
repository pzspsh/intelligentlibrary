/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:05:36
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := getHostIp()
	fmt.Println(ip)
}

func getHostIp() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get current host ip err: ", err)
		return ""
	}
	var ip string
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ip = ipNet.IP.String()
				break
			}
		}
	}
	return ip
}
