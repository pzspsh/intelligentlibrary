/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:23:49
*/
package main

import (
	"fmt"
	"net"
	"syscall"
)

func main() {
	// 创建 socket
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer syscall.Close(fd)

	// 连接服务器
	addr := syscall.SockaddrInet4{Port: 8000}
	copy(addr.Addr[:], net.ParseIP("localhost").To4())
	err = syscall.Connect(fd, &addr)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}

	// 发送数据
	message := "Hello, server!"
	_, err = syscall.Write(fd, []byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	// 接收数据
	buffer := make([]byte, 1024)
	n, err := syscall.Read(fd, buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}
	fmt.Println("Received message:", string(buffer[:n]))
}
