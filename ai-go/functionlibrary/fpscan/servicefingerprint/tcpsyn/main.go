/*
@File   : main.go
@Author : pan
@Time   : 2024-06-11 11:17:19
*/
package main

/*
对于非HTTP服务，我们可以使用Go的net包来发送原始的TCP SYN包，并观察响应。这种方法更底层，通常用于分析TCP/IP堆栈的行为差异。
*/
import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := fmt.Sprintf("%s:%v", "target_ip_address", "target_port") // 替换为目标IP和端口，例如SSH端口

	conn, err := net.DialTimeout("tcp", target, 5*time.Second) // 设置连接超时
	if err != nil {
		fmt.Printf("Error connecting to %s: %v\n", target, err)
		return
	}
	defer conn.Close()

	// 在此处可以发送特定于服务的请求，并分析响应
	// 但对于大多数服务，简单的连接尝试可能就足够了

	fmt.Printf("Successfully connected to %s\n", target)
}
