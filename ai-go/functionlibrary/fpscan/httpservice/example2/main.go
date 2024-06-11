/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 10:37:36
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func detectHTTP(host string, port int) bool {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("http://%s:%d/", host, port))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return resp.StatusCode == http.StatusOK
	}
	return false
}

func detectHTTPS(host string, port int) bool {
	// 创建一个配置为不验证TLS证书的HTTP客户端
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}
	resp, err := client.Get(fmt.Sprintf("https://%s:%d/", host, port))
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	// 读取响应体，以确保连接是完整的
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	if resp.StatusCode == http.StatusOK {
		return true
	}
	return false
}

func main() {
	host := "127.0.0.1" // 替换为你想要探测的IP或主机名
	port := 7001        // 替换为你想要探测的端口

	if detectHTTP(host, port) {
		fmt.Printf("Port %d on %s supports HTTP\n", port, host)
	} else {
		fmt.Printf("Port %d on %s does not support HTTP\n", port, host)
	}

	if detectHTTPS(host, port) {
		fmt.Printf("Port %d on %s supports HTTPS\n", port, host)
	} else {
		fmt.Printf("Port %d on %s does not support HTTPS\n", port, host)
	}
}
