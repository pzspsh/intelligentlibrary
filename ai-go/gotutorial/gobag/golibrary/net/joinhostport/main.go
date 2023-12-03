/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 12:41:20
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr) //[127.0.0.1/8 10.236.15.24/24 ::1/128 fe80::3617:ebff:febe:f123/64]
	hp := net.JoinHostPort("127.0.0.1", "8080")
	fmt.Println(hp) //127.0.0.1:8080,根据ip和端口组成一个addr 字符串 表示
	lt, _ := net.LookupAddr("127.0.0.1")
	fmt.Println(lt) //[ localhost ],根据地址查找到改地址的一个映射列表
	host, _ := net.LookupHost("www.baidu.com")
	fmt.Println(host) //[111.13.100.92 111.13.100.91],查找给定域名的host名称
}
