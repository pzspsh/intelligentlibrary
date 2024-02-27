/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:38:18
*/
package main

import (
	"context"
	"fmt"
	"net"
	"time"
)

func main() {
	//创建一个context对象，用于超时控制
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//设置Google DNS服务器的地址
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}

	//解析域名
	ips, err := resolver.LookupIPAddr(ctx, "google.com")
	if err != nil {
		panic(err)
	}

	//打印解析结果
	for _, ip := range ips {
		fmt.Println(ip.IP)
	}
}
