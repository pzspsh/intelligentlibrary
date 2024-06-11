/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:53:32
*/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := fmt.Sprintf("%s:%d", "example.com", 22) // 替换为目标IP和SSH端口

	conn, err := net.DialTimeout("tcp", target, 5*time.Second)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", target, err)
		return
	}
	defer conn.Close()

	// 发送SSH版本标识字符串
	version := []byte("SSH-2.0-Go_SSH_Fingerprint_Example\r\n")
	_, err = conn.Write(version)
	if err != nil {
		fmt.Printf("Error sending SSH version: %v\n", err)
		return
	}

	// 读取响应
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading SSH response: %v\n", err)
		return
	}

	// 分析响应（这里只是简单输出）
	fmt.Printf("SSH Response: %s\n", buffer[:n])
}
