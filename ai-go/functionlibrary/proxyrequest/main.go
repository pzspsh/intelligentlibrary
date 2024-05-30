/*
@File   : main.go
@Author : pan
@Time   : 2024-05-30 17:27:31
*/
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	proxyUrl, err := url.Parse("http://127.0.0.1:8080") // http://代理服务器IP:端口
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误解析代理URL: %v\n", err)
		os.Exit(1)
	}
	proxyclient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	targetUrl := "http://www.example.com"   // 请求目标
	resp, err := proxyclient.Get(targetUrl) // 发送请求
	if err != nil {
		fmt.Fprintf(os.Stderr, "错误发送请求: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Println("响应状态码:", resp.StatusCode)
}
