/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:27
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1.与Server端建立TCP链接
	address := "10.20.172.108:22022"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("Connect Server failed! \n [err]: %v\n", err)
		return
	} else {
		fmt.Printf("Connect Server %v successful!\n", address)
	}
	// 退出通信连接
	defer conn.Close()

	// 2.发送初始连接信息到Server端
	sendMsg := fmt.Sprintf("Hello Server, I'm  %v client.", conn.LocalAddr())
	inputReader := bufio.NewReader(os.Stdin) // 复习点
	conn.Write([]byte(sendMsg))

	// 3.循环从服务端接收以及从终端输入发送信息到服务端
	for {
		// 服务端回复的信息
		reply := [1024]byte{}
		n, err := conn.Read(reply[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Printf("Server > %v\n", string(reply[:n]))

		// 客户端输入字符串信息
		fmt.Printf("请输入消息:")
		sendMsg, _ = inputReader.ReadString('\n') // 复习点 (以\n截至读取)
		sendMsg = strings.TrimSpace(sendMsg)      // 复习点 处理输入字符串的前后空格
		sendMsg = strings.Trim(sendMsg, "\n")     // 复习点 处理输入字符串的最后的换行符

		// 但客户端输入quit则退出与服务端建立的TCP通信连接.
		if strings.ToUpper(sendMsg) == "QUIT" {
			fmt.Printf("exit conn.......")
			break
		}

		// 发送已经处理过后的信息到服务端
		_, err = conn.Write([]byte(sendMsg))
		if err != nil {
			return
		}
	}
}
