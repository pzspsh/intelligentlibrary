/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:51:18
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr := "127.0.0.1:8888"
	localAddr := "127.0.0.1:0"

	// 解析服务器地址
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		fmt.Println("Error resolving server address:", err)
		return
	}

	// 解析本地地址
	localUdpAddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		fmt.Println("Error resolving local address:", err)
		return
	}

	// 创建UDP socket
	conn, err := net.DialUDP("udp", localUdpAddr, udpAddr)
	if err != nil {
		fmt.Println("Error dialing server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", serverAddr)

	// 发送数据
	message := []byte("Hello, server!")
	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending:", err)
		return
	}

	// 接收数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received message:", string(buffer[:n]))
}
