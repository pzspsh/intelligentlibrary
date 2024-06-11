/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:44:45
*/
package main

import (
	"fmt"
	"net"
	"time"
)

/*
通过发送特定的TCP数据包并观察响应，可以分析远程主机的TCP/IP堆栈行为。
不同的操作系统和配置可能会导致不同的TCP选项、窗口大小、MSS值等。
虽然这种方法需要深入的网络知识，但它可以提供比基于应用层协议的分析更底层的线索。
*/

func main() {
	target := "example.com:80" // 替换为目标IP和端口

	// 设置连接超时
	conn, err := net.DialTimeout("tcp", target, 5*time.Second)
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", target, err)
		return
	}
	defer conn.Close()

	// 发送TCP SYN包（这一步通常由操作系统自动处理）
	// 接下来可以发送特定的TCP选项或进行其他TCP交互

	// 读取并分析TCP响应（这通常需要更底层的socket编程知识）
	// 分析TCP选项、窗口大小、MSS等，但这通常需要对TCP协议有深入的了解

	// 此时，TCP连接的SYN包和SYN-ACK包交换已经完成，连接已建立
	// 我们可以通过conn对象来发送和接收数据

	// 发送一些数据（可选）
	_, err = conn.Write([]byte("GET / HTTP/1.1\r\nHost: example.com\r\n\r\n"))
	if err != nil {
		fmt.Printf("Error sending data: %v\n", err)
		return
	}

	// 读取响应数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		return
	}

	fmt.Printf("Received %d bytes: %s\n", n, buffer[:n])

	/*
		建立了一个到目标IP和端口的TCP连接，并发送了一个简单的HTTP GET请求。然后，我们尝试读取服务器的响应。
		这个例子并不直接观察TCP堆栈行为，但它展示了如何使用Go与TCP服务进行交互。

		要直接观察TCP堆栈行为（如SYN包和ACK包的交换），您通常需要使用更底层的工具，如tcpdump（在命令行中使用）
		或专门的库（可能需要在C或其他底层语言中实现），这些工具可以捕获和分析网络数据包。这些工具允许您观察TCP
		连接建立过程中的每个数据包，包括SYN包、SYN-ACK包和ACK包，以及它们的选项和负载。
	*/
}
