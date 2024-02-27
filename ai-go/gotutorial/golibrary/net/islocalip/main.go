/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 12:44:07
*/
package main

import "net"

func isLocalIP(ip string) (bool, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false, err
	}

	for i := range addrs {
		intf, _, err := net.ParseCIDR(addrs[i].String()) //net.ParseCIDR 解析IP
		if err != nil {
			return false, err
		}
		if net.ParseIP(ip).Equal(intf) {
			return true, nil
		}
	}
	return false, nil
}

func main() {
	isLocalIP("")
}
