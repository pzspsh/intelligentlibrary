/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:57:06
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	// 创建TCP拨号器并设置超时时间
	dialer := &net.Dialer{
		Timeout: 30 * time.Second,
	}
	// 创建TLS配置和自定义拨号器
	tlsConfig := &tls.Config{}
	tlsDialer := &tls.Dialer{
		NetDialer: dialer,
		Config:    tlsConfig,
	}

	// 发送https请求并读取响应内容
	conn, err := tlsDialer.Dial("tcp", "www.example.com:443")
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}

	request := []byte("GET / HTTP/1.0\r\nHost: www.example.com\r\n\r\n")
	_, err = conn.Write(request)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}

	response, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}
	fmt.Println(string(response))
}
