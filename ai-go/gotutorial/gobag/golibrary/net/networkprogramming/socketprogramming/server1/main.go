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

	// 绑定地址
	addr := syscall.SockaddrInet4{Port: 8000}
	copy(addr.Addr[:], net.ParseIP("localhost").To4())
	err = syscall.Bind(fd, &addr)
	if err != nil {
		fmt.Println("Error binding address:", err)
		return
	}

	// 监听连接请求
	err = syscall.Listen(fd, syscall.SOMAXCONN)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	fmt.Println("Listening on :8000")

	for {
		// 接受连接请求
		connFd, _, err := syscall.Accept(fd)
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		defer syscall.Close(connFd)

		// 处理连接
		go handleConnection(connFd)
	}
}

func handleConnection(fd syscall.Handle) {
	// 接收数据
	buffer := make([]byte, 1024)
	n, err := syscall.Read(fd, buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}
	fmt.Println("Received message:", string(buffer[:n]))

	// 发送数据
	message := "Hello, client!"
	_, err = syscall.Write(fd, []byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
}
