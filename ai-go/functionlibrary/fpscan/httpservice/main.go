/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 10:10:49
*/
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// isHTTPService 判断指定端口上的服务是否是HTTP服务
func isHTTPService(host string, port int) (bool, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 5*time.Second) // 设置连接超时
	if err != nil {
		return false, err
	}
	defer conn.Close()

	// 尝试发送HTTP GET请求
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d", host, port), nil)
	if err != nil {
		return false, err
	}

	// 使用http.Client发送请求，但禁用自动重定向和TLS验证
	client := &http.Client{
		Transport: &http.Transport{
			DialContext:       func(_ context.Context, _, _ string) (net.Conn, error) { return conn, nil },
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true}, // 仅为HTTPS测试，实际中不建议禁用TLS验证
			ForceAttemptHTTP2: false,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求失败: ", err) // 如果连接被拒绝或请求失败，可能不是HTTP服务
		return false, err
	}
	defer resp.Body.Close()

	// 读取响应体，这里只是简单地读取并丢弃，因为我们只关心响应头
	_, _ = io.ReadAll(resp.Body)

	// 检查响应头中是否包含HTTP相关的标识
	fmt.Println(resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	return false, err
}

// isHTTPSService 判断指定端口上的服务是否是HTTPS服务
func isHTTPSService(host string, port int) (bool, error) {
	// 尝试使用TLS握手来检测HTTPS服务
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("TLS握手失败: ", err)
		return false, err
	}
	defer conn.Close()

	// 如果TLS握手成功，则可能是HTTPS服务
	// 注意：这里只是一个简单的检测，因为非HTTPS服务也可能实现TLS握手
	return true, nil
}

func main() {
	host := "127.0.0.1"
	httpPort := 7001
	httpsPort := 10000

	isHTTP, err := isHTTPService(host, httpPort)
	if err != nil {
		fmt.Printf("Error checking HTTP service: %s\n", err)
	} else {
		if isHTTP {
			fmt.Printf("The service running on port %d of %s is HTTP.\n", httpPort, host)
		} else {
			fmt.Printf("The service running on port %d of %s is not HTTP.\n", httpPort, host)
		}
	}

	isHTTPS, err := isHTTPSService(host, httpsPort)
	if err != nil {
		fmt.Printf("Error checking HTTPS service: %s\n", err)
	} else {
		if isHTTPS {
			fmt.Printf("The service running on port %d of %s is HTTPS.\n", httpsPort, host)
		} else {
			fmt.Printf("The service running on port %d of %s is not HTTPS.\n", httpsPort, host)
		}
	}
}
