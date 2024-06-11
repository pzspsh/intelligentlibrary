/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 10:36:45
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// probeHTTPPort 尝试连接到指定端口的HTTP服务，并检查是否返回HTTP响应
func probeHTTPPort(host string, port int, timeout time.Duration) (bool, error) {
	target := fmt.Sprintf("http://%s:%d", host, port)
	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(target)
	if err != nil {
		return false, fmt.Errorf("failed to probe HTTP on %s:%d: %w", host, port, err)
	}
	defer resp.Body.Close()

	// 读取并丢弃响应体内容
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read HTTP response body on %s:%d: %w", host, port, err)
	}

	// 只要收到HTTP响应，就认为端口支持HTTP，无论状态码是什么
	fmt.Printf("Port %d on %s supports HTTP. Status Code: %d\n", port, host, resp.StatusCode)
	return true, nil
}

func main() {
	host := "127.0.0.1"
	port := 7001
	timeout := 5 * time.Second

	// 探测HTTP端口
	supportsHTTP, err := probeHTTPPort(host, port, timeout)
	if err != nil {
		fmt.Printf("Error probing HTTP port %d on %s: %v\n", port, host, err)
	} else if supportsHTTP {
		fmt.Printf("Port %d on %s supports HTTP\n", port, host)
	} else {
		fmt.Printf("Unexpected error: Function should return true if an HTTP response is received.\n")
	}
}
