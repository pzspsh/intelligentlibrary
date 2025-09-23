package main

import (
	"fmt"
	"net"
)

func main() {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网卡列表失败：", err)
		return
	}
	for _, iface := range ifaces {
		fmt.Println(iface.Name, iface.HardwareAddr)
	}
}
