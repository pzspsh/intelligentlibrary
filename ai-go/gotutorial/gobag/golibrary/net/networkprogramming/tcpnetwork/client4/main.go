/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:27
*/
package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	message       = "Ping"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int) {
	// 根据端口拼接网络地址
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	// 根据地址拨号
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer conn.Close()
	// 写入发送消息
	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	// 打印发送消息
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	// 循环读取消息，响应服务器
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {

	var (
		ip   = "127.0.0.1"
		port = 3333
	)
	// 拨号、建立与服务端连接
	SocketClient(ip, port)

}
