/*
@File   : main.go
@Author : pan
@Time   : 2024-12-17 09:47:54
*/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 目标IP地址和端口号
	ip := "10.0.35.64"
	port := "22"

	// 连接到目标
	conn, err := net.DialTimeout("tcp", ip+":"+port, 10*time.Second)
	if err != nil {
		fmt.Println("无法连接到目标:", err)
		return
	}
	defer conn.Close()

	// 设置读取超时时间
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	// 读取数据
	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("读取数据出错:", err)
		return
	}

	// 解析操作系统信息
	os := parseOS(buffer)
	fmt.Printf("操作系统: %s\n", os)
}

func parseOS(data []byte) string {
	// 这里是一个简单的示例，根据TCP/IP选项来猜测操作系统
	// 实际应用中可能需要更复杂的逻辑和更多的数据来提高准确性
	fmt.Println(string(data))
	if data[0] == 240 && data[1] == 255 {
		return "Linux"
	} else if data[0] == 255 && data[1] == 255 {
		return "Windows"
	} else {
		return "Unknown"
	}
}
