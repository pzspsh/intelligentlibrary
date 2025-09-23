package main

import (
	"fmt"
	"net"
)

func main() {
	ip := net.ParseIP("192.168.0.1") // 解析IP字符串
	addr := net.IPAddr{
		IP:   ip,
		Zone: "", // 可选的接口域（zone）
	}

	fmt.Println(addr)
}
