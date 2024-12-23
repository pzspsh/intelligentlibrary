/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:05:56
*/
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	// 1.服务端开启监听 UDP 通信的相关设置
	server_ip := [4]byte{10, 20, 172, 108}
	server_port := 30000
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(server_ip[0], server_ip[1], server_ip[2], server_ip[3]),
		Port: server_port,
	})

	if err != nil {
		fmt.Printf("Listen UDP Server (%v:%v) Failed, err: %v\n", server_ip, server_port, err)
		return
	} else {
		fmt.Printf("[%v] Listening UDP Server %v:%v is successful!\n", time.Now().Format("2006-01-02 15:04:05"), server_ip, server_port)
	}
	// 程序结束时关闭conn资源
	defer conn.Close()

	// 2.循环接收和响应数据给客户端，非常主要此处不需要建立连接，直接收发数据。
	for {
		// 获得客户端通信对象以及返回读取的字节数
		var recvMsg [1024]byte
		count, addr, err := conn.ReadFromUDP(recvMsg[:]) // 接收数据
		if err != nil {
			fmt.Println("Read from UDP Client failed. Err:", err)
			return
		}
		// 打印客户端发送的信息到终端
		fmt.Printf("[%v] %v - %v\n", time.Now().Format("2006-01-02 15:04:05"), addr.String(), string(recvMsg[:count-1]))

		// 并将接收到的信息更改为大写，再返还给Client。
		reply := strings.ToUpper(string(recvMsg[:count]))
		conn.WriteToUDP([]byte(reply), addr) // 发送数据
	}
}
