/*
@File   : main.go
@Author : pan
@Time   : 2024-06-07 17:19:18
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func isMySQLService(host string, port int) (bool, error) {
	// 设置连接超时
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 5*time.Second)
	if err != nil {
		return false, err
	}
	defer conn.Close()

	// MySQL协议握手消息的前4个字节
	handshake := []byte{0x0a, 0x00, 0x00, 0x01}

	// 发送握手消息
	if _, err := conn.Write(handshake); err != nil {
		return false, err
	}

	// 读取响应，MySQL协议握手响应通常会比较长，这里我们只读取一小部分作为示例
	reader := bufio.NewReader(conn)
	response := make([]byte, 10)
	if _, err := reader.Read(response); err != nil {
		return false, err
	}
	fmt.Println(response)
	fmt.Printf("Response: %v\n", string(response))
	// 根据响应判断是否是MySQL服务，这里只是一个简化的检查，真实情况可能更复杂
	// 注意：这个方法并不完全可靠，因为有些非MySQL服务可能也会响应类似的字节序列
	// 更好的方法是检查响应中更具体的MySQL协议特征
	fmt.Println(response[0], response[1], response[2], response[3])
	isMySQL := response[0] == 0x0a && response[1] == 0x57 && response[2] >= 0x08 && response[3] == 0x00
	return isMySQL, nil
}

func main() {
	host := "127.0.0.1" // 替换成你想要探测的主机地址
	port := 3306        // MySQL默认端口

	isMySQL, err := isMySQLService(host, port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error detecting MySQL service: %s\n", err)
		os.Exit(1)
	}

	if isMySQL {
		fmt.Printf("The service running on port %d of %s is MySQL.\n", port, host)
	} else {
		fmt.Printf("The service running on port %d of %s is not MySQL.\n", port, host)
	}
}
