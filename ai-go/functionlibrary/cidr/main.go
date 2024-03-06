/*
@File   : main.go
@Author : pan
@Time   : 2024-03-06 10:51:24
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("116.198.164.83")
	mask := net.CIDRMask(24, 32) // 这里的 24 表示子网掩码的位数，32 是 IPv4 地址的总位数
	network := ip.Mask(mask)
	ipNet := &net.IPNet{
		IP:   network,
		Mask: mask,
	}
	fmt.Println(ipNet)
}
