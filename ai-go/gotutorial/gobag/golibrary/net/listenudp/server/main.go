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
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":22333") // 转换地址，作为服务器使用时需要监听本机的一个端口
	// 端口号写 0 可以由系统随机分配可用的端口号
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr) // 启动UDP监听本机端口
	checkError(err)

	for {
		var buf [128]byte
		len, addr, err := conn.ReadFromUDP(buf[:]) // 读取数据，返回值依次为读取数据长度、远端地址、错误信息 // 读取操作会阻塞直至有数据可读取
		checkError(err)
		fmt.Println(string(buf[:len])) // 向终端打印收到的消息

		_, err = conn.WriteToUDP([]byte("233~~~"), addr) // 写数据，返回值依次为写入数据长度、错误信息 // WriteToUDP()并非只能用于应答的，只要有个远程地址可以随时发消息
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
