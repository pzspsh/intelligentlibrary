/*
@File   : main.go
@Author : pan
@Time   : 2024-05-30 17:40:28
*/
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 获取系统代理设置
	proxyURL, err := http.ProxyFromEnvironment(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting system proxy: %v\n", err)
		os.Exit(1)
	}

	// 创建代理传输设置
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	// 创建使用代理的客户端
	client := &http.Client{
		Transport: transport,
	}

	// 请求的目标URL
	targetURL := "http://www.example.com"

	// 发送GET请求
	resp, err := client.Get(targetURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 输出响应头
	for name, values := range resp.Header {
		fmt.Printf("%s: %s\n", name, values)
	}
}
