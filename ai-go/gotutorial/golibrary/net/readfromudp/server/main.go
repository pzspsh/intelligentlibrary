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
	// 监听IP地址和端口号
	addr := net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8888,
	}

	// 创建UDP socket
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Server is listening on", addr.String())

	// 读取数据
	buffer := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Println("Received message:", string(buffer[:n]), "from", remoteAddr)

	// 发送数据
	message := []byte("Hello, client!")
	_, err = conn.WriteToUDP(message, remoteAddr)
	if err != nil {
		fmt.Println("Error sending:", err)
		return
	}

	fmt.Println("Message sent")
}
