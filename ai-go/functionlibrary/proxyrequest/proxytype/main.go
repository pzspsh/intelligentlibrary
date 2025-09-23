package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// 判断代理类型
func detectProxyType(proxyAddr string) (string, error) {
	// 尝试解析代理地址
	if strings.HasPrefix(proxyAddr, "http://") || strings.HasPrefix(proxyAddr, "https://") {
		return "HTTP/HTTPS", nil
	} else if strings.HasPrefix(proxyAddr, "socks5://") {
		return "SOCKS5", nil
	}

	// 去掉可能存在的协议前缀
	proxyAddr = strings.TrimPrefix(proxyAddr, "http://")
	proxyAddr = strings.TrimPrefix(proxyAddr, "https://")
	proxyAddr = strings.TrimPrefix(proxyAddr, "socks5://")

	// 分割地址和端口
	host, port, err := net.SplitHostPort(proxyAddr)
	if err != nil {
		return "", fmt.Errorf("无效的代理地址: %v", err)
	}

	// 默认端口
	if port == "" {
		port = "8080" // 默认 HTTP 代理端口
	}

	// 尝试检测代理类型
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 5*time.Second)
	if err != nil {
		return "", fmt.Errorf("无法连接到代理服务器: %v", err)
	}
	defer conn.Close()

	// 设置超时时间
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	// 尝试发送 SOCKS5 握手请求
	socks5Response, err := trySOCKS5Handshake(conn)
	if err == nil && socks5Response {
		return "SOCKS5", nil
	}

	// 尝试发送 HTTP 请求
	httpResponse, err := tryHTTPRequest(conn)
	if err == nil && httpResponse {
		return "HTTP/HTTPS", nil
	}

	return "未知类型", nil
}

// 尝试 SOCKS5 握手
func trySOCKS5Handshake(conn net.Conn) (bool, error) {
	// 发送 SOCKS5 握手请求
	_, err := conn.Write([]byte{0x05, 0x01, 0x00}) // 协议版本 5，1 种认证方法，无认证
	if err != nil {
		return false, err
	}

	// 读取响应
	buf := make([]byte, 2)
	_, err = bufio.NewReader(conn).Read(buf)
	if err != nil {
		return false, err
	}

	// 检查响应是否有效
	return buf[0] == 0x05 && buf[1] == 0x00, nil
}

// 尝试 HTTP 请求
func tryHTTPRequest(conn net.Conn) (bool, error) {
	// 发送简单的 HTTP 请求
	request := "GET http://example.com HTTP/1.1\r\nHost: example.com\r\n\r\n"
	_, err := conn.Write([]byte(request))
	if err != nil {
		return false, err
	}

	// 读取响应
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	// 检查响应是否是 HTTP 响应
	return strings.HasPrefix(line, "HTTP/"), nil
}

func main() {
	// 测试代理地址
	proxies := []string{
		"http://127.0.0.1:8080",
		"socks5://127.0.0.1:1080",
		"127.0.0.1:8080", // 未知类型
	}

	for _, proxy := range proxies {
		proxyType, err := detectProxyType(proxy)
		if err != nil {
			fmt.Printf("检测代理失败 (%s): %v\n", proxy, err)
		} else {
			fmt.Printf("代理类型 (%s): %s\n", proxy, proxyType)
		}
	}
}
