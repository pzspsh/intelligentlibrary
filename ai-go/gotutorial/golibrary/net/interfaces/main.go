/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 14:22:35
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	interfaces, _ := net.Interfaces()
	for _, addr := range interfaces {
		fmt.Println(addr.Name)         // 物理网卡设备名
		fmt.Println(addr.Index)        //索引
		fmt.Println(addr.HardwareAddr) // mac地址
		fmt.Println(addr.Flags)
		fmt.Println(addr.MTU)              // 组大传输数据
		fmt.Println(addr.MulticastAddrs()) // 组播地址列表
	}
}
