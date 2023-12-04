/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 14:06:06
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 从环境变量中获取代理信息
	proxy := http.ProxyFromEnvironment

	// 创建HTTP客户端
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
		},
	}

	// 发送GET请求
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
		// 处理错误
	}

	// 关闭响应体
	defer resp.Body.Close()

	// 读取响应内容
	// ...
}
