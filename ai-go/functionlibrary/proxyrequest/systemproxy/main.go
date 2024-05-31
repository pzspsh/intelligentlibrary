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
	proxyURL, err := http.ProxyFromEnvironment(nil) // 获取系统代理设置
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting system proxy: %v\n", err)
		os.Exit(1)
	}
	transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)} // 创建代理传输设置
	client := &http.Client{Transport: transport} // 创建使用代理的客户端
	targetURL := "http://www.example.com" // 请求的目标URL
	resp, err := client.Get(targetURL) // 发送GET请求
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	for name, values := range resp.Header { // 输出响应头
		fmt.Printf("%s: %s\n", name, values)
	}
}
