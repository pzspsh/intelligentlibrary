/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:43:24
*/
package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", "192.168.31.189:53771") // 转换地址，作为客户端使用要向远程发送消息，这里用远程地址与端口号
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr) // 建立连接，第二个参数为nil时通过默认本地地址（猜测可能是第一个可用的地址，未进行测试）发送且端口号自动分配，第三个参数为远程端地址与端口号
	checkError(err)

	go receive(conn) // 使用DialUDP建立连接后也可以监听来自远程端的数据

	for {
		_, err = conn.Write([]byte("naisu233~~~")) // 向远程端发送消息
		checkError(err)
		time.Sleep(4 * time.Second) // 等待4s
	}
}

func receive(conn *net.UDPConn) {
	for {
		var buf [128]byte
		len, err := conn.Read(buf[0:]) // 读取数据 // 读取操作会阻塞直至有数据可读取
		checkError(err)
		fmt.Println(string(buf[0:len]))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
