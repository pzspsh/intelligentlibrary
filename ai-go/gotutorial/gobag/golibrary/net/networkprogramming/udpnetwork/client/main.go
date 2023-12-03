/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:05:56
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// (1) 与服务端建立UDP通信链接
	server_ip := [4]byte{10, 20, 172, 108}
	server_port := 30000
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(server_ip[0], server_ip[1], server_ip[2], server_ip[3]),
		Port: server_port,
	})
	if err != nil {
		fmt.Printf("Connect UDP Server (%v:%v) Failed! err: %v\n", server_ip, server_port, err)
		return
	} else {
		fmt.Printf("[%v] - Connect UDP Server %v:%v successful!\n", time.Now().Format("2006-01-02 15:04:05"), server_ip, server_port)
	}
	// 同样关闭建立的通信连接
	defer socket.Close()

	// (2) 发送与接收服务端返回的信息
	var reply [1024]byte
	inputData := bufio.NewReader(os.Stdin)
	for {
		// 终端接收输入要发送的给服务端的内容
		fmt.Print("<- 请输入将要发送的内容:")
		sendMsg, _ := inputData.ReadString('\n')
		sendMsg = strings.TrimSpace(sendMsg) // 取消字符串前后的空格
		socket.Write([]byte(sendMsg))        // 发送数据
		if err != nil {
			fmt.Printf("发送数据失败，err: %v\n", err)
			return
		}
		// 接收来自服务端的反馈的内容
		count, _, err := socket.ReadFromUDP(reply[:]) // 接收数据
		if err != nil {
			fmt.Printf("接收数据失败, err: %v\n", err)
			return
		}
		fmt.Printf("Server -> [%v Bytes] %v \n", count, string(reply[:count]))
	}
}
