/*
@File   : main.go
@Author : pan
@Time   : 2024-11-27 14:40:40
*/
package main

import (
	"fmt"
	"net"
	"time"
)

// 尝试连接HTTP代理
func detectHTTPProxy(ip string, port int) bool {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	// 发送CONNECT命令
	_, err = conn.Write([]byte("CONNECT / HTTP/1.1\r\nHost: example.com\r\n\r\n"))
	if err != nil {
		return false
	}

	// 读取响应的前几个字节判断是否成功
	buf := make([]byte, 10)
	n, err := conn.Read(buf)
	if err != nil || n < 10 {
		return false
	}

	return string(buf[:10]) == "HTTP/1.1 2"
}

// 尝试连接HTTPS代理
func detectHTTPSProxy(ip string, port int) bool {
	// HTTPS代理通常也支持HTTP代理的CONNECT方法
	return detectHTTPProxy(ip, port)
}

// 尝试连接SOCKS5代理
func detectSOCKS5Proxy(ip string, port int) bool {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()

	// 发送SOCKS5握手包和请求包
	_, err = conn.Write([]byte{0x05, 0x01, 0x00}) // 握手包
	if err != nil {
		return false
	}
	buf := make([]byte, 2)
	n, err := conn.Read(buf)
	if err != nil || n != 2 || buf[0] != 0x05 || buf[1] != 0x00 {
		return false
	}

	_, err = conn.Write([]byte{0x05, 0x01, 0x00, 0x03, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x00, 0x50, 0x80}) // 请求包
	if err != nil {
		return false
	}
	buf = make([]byte, 10)
	n, err = conn.Read(buf)
	if err != nil || n != 10 {
		return false
	}

	return string(buf[:3]) == "05 00"
}

func main() {
	ip := "your_proxy_ip"
	port := 8080

	isHTTP := detectHTTPProxy(ip, port)
	isHTTPS := detectHTTPSProxy(ip, port)
	isSOCKS5 := detectSOCKS5Proxy(ip, port)

	if isHTTP {
		fmt.Println("Proxy is HTTP")
	}
	if isHTTPS {
		fmt.Println("Proxy is HTTPS")
	}
	if isSOCKS5 {
		fmt.Println("Proxy is SOCKS5")
	}
}
