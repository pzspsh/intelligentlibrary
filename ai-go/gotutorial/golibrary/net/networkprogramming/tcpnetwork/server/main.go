/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:45
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// 建立socket监听
	lis, err := net.Listen("tcp", "localhost:1024")
	log.Println("server up at " + lis.Addr().String())
	defer lis.Close()
	log.Println("Waiting for clients...")

	// 处理错误
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}

	// 处理客户端连接
	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}

		log.Printf("client addr: %s, tcp connect success", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	buffer := make([]byte, 2048)
	// 循环读取客户请求
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("%s connection error: %s", conn.RemoteAddr(), err)
			return
		}
		log.Printf("From %s receive data string: %s\n", conn.RemoteAddr(), string(buffer[:n]))
		// 收到的返回信息
		strTmp := fmt.Sprintf("server got msg \"%s\" at %s", string(buffer[:n]), time.Now().String())
		conn.Write([]byte(strTmp))
	}
}
