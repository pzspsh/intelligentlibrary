/*
@File   : main.go
@Author : pan
@Time   : 2024-05-30 17:37:06
*/
package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/proxy"
)

func main() {
	// 创建一个代理Dialer
	dialer, err := proxy.SOCKS5("tcp", "proxy_address:port", nil, proxy.Direct)
	if err != nil {
		panic(err)
	}

	// 使用代理Dialer创建一个自定义的Transport
	transport := &http.Transport{
		Dial: dialer.Dial, // 使用代理Dialer代替默认的Dialer
	}

	// 使用自定义Transport创建一个Client
	client := &http.Client{
		Transport: transport,
	}

	// 使用Client发起HTTP请求
	resp, err := client.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应内容并打印
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
