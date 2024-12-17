/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 11:16:19
*/
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	useNetParse()
}

func useNetParse() {
	host := "www.baidu.com"
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range ips {
		fmt.Println("DNS 解析：", v.String())
	}

	ip, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务器 IP 地址：", ip.String())
}
