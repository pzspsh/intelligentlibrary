/*
@File   : main.go
@Author : pan
@Time   : 2024-08-01 12:07:17
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// 代理服务器的URL
	proxyURL, _ := url.Parse("https://proxy.example.com:8080")

	// 如果代理需要认证，可以在URL中添加用户名和密码
	// proxyURL.User = url.UserPassword("username", "password")

	// 配置代理
	http.DefaultTransport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)

	// 如果你的代理使用了自签名证书，你需要忽略证书验证
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	// 创建一个HTTP客户端
	client := &http.Client{}

	// 要请求的目标网站URL
	targetURL := "https://www.example.com"

	// 创建一个HTTP请求
	req, _ := http.NewRequest("GET", targetURL, nil)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
