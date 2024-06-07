/*
@File   : main.go
@Author : pan
@Time   : 2024-06-07 15:09:18
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func detectSSHService(target string, port string) (string, error) {
	// 连接到目标端口
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", target, port))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// 创建一个读取器来读取响应
	reader := bufio.NewReader(conn)

	// SSH协议通常在连接建立后会发送一个版本字符串
	// 尝试读取并解析这个字符串
	banner, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	fmt.Println(banner)
	// 清理字符串，移除换行符
	banner = strings.TrimSuffix(banner, "\n")

	// 检查是否包含SSH的标识字符串
	if strings.Contains(banner, "SSH") {
		return "SSH", nil
	}

	// 如果不包含SSH标识，返回未知服务
	return "unknown", nil
}

func main() {
	target := "127.0.0.1" // 替换为目标IP或域名
	port := "22"          // 端口号

	service, err := detectSSHService(target, port)
	if err != nil {
		fmt.Printf("Error detecting service: %s\n", err)
		return
	}

	fmt.Printf("Detected service on port %s: %s\n", port, service)
}
