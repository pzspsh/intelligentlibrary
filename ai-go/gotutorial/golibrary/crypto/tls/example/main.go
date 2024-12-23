/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 12:26:41
*/
package main

import (
	"crypto/tls"
	"io"
	"log"
	"net"
)

func main() {
	// 创建一个监听器
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	// 加载证书和私钥
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
	// 配置 TLS 连接
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// 使用 TLS 封装连接
		tlsConn := tls.Server(conn, &config)
		// 开启一个协程处理连接
		go handleConnection(tlsConn)
	}
}
func handleConnection(conn net.Conn) {
	// 建立目标服务器连接
	serverConn, err := net.Dial("tcp", "example.com:443")
	if err != nil {
		log.Fatal(err)
	}
	defer serverConn.Close()
	// 将客户端发送的数据转发到目标服务器
	go func() {
		_, err := io.Copy(serverConn, conn)
		if err != nil {
			log.Fatal(err)
		}
	}()
	// 将目标服务器响应的数据转发回客户端
	_, err = io.Copy(conn, serverConn)
	if err != nil {
		log.Fatal(err)
	}
}
