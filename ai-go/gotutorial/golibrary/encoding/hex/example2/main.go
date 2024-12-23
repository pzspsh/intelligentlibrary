/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:41:50
*/
package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"os"
)

func main() {
	// 建立 TCP 连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}

	// 将要发送的 16 进制数据转换为字节数组
	hexdata := "68656c6c6f"
	data, err := hex.DecodeString(hexdata)
	if err != nil {
		fmt.Println("Hex decoding error:", err.Error())
		os.Exit(1)
	}

	// 发送数据
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		os.Exit(1)
	}

	// 接收数据
	reply := make([]byte, 1024)
	n, err := conn.Read(reply)
	if err != nil {
		fmt.Println("Error receiving:", err.Error())
		os.Exit(1)
	}

	// 输出服务端回复的内容
	fmt.Println("Server reply:", string(reply[:n]))

	// 关闭连接
	conn.Close()
}
