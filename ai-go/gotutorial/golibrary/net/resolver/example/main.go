/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:40:02
*/
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func main() {
	// 创建一个新的 Dialer
	d := &net.Dialer{
		Timeout:   30 * time.Second, // 连接超时时间
		KeepAlive: 30 * time.Second, // 保持连接
		DualStack: true,             // 支持 IPv4 和 IPv6
	}
	// 设置 DNS
	resolver := &net.Resolver{
		PreferGo: true,
		Dial:     d.DialContext,
	}

	// 将代理设置为 http.Transport 中的 Dial 函数
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 连接超时时间
			KeepAlive: 30 * time.Second, // 保持连接
			DualStack: true,             // 支持 IPv4 和 IPv6
			Resolver:  resolver,         // 使用新设置的解析器
		}).DialContext,
		TLSHandshakeTimeout: 10 * time.Second, // TLS 握手超时时间
	}

	// 设置 http 客户端
	client := &http.Client{
		Timeout:   time.Second * 60, // 超时时间
		Transport: transport,        // 使用新设置的 transport
	}

	// 访问一个带有 DNS 规则的网站
	req, err := http.NewRequest(http.MethodGet, "http://www.google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
