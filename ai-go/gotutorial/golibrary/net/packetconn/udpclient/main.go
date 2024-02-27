/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:19:56
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	message = "网友消息"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:6666")
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	if err != nil {
		log.Fatalf("Dial is error %v", err)
		os.Exit(1) // 退出
	}
	conn.Write([]byte(message))
	log.Printf("Send: %s", message)

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalf("Read is error", n)
	}
	log.Printf("Received： 客户端接收信息： %s", buf[:n])
}
