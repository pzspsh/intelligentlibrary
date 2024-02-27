/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:45
*/
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	Message       = "Pong"
	StopCharacter = "\r\n\r\n"
)

func SocketServer(port int) {
	// 向服务器发起tcp协议的网络启动连接
	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))

	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	//
	// 退出函数前，把返回来的struct结构体方法关闭
	// type Listener interface {
	//		Accept() (Conn, error)
	//		Close() error
	//		Addr() Addr
	//	}
	defer listen.Close()

	log.Printf("开始监听端口号: %d", port)

	for {
		// 接收输入流
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		// 解析输入流
		go handler(conn)
	}

}

func handler(conn net.Conn) {

	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			if isTransportOver(data) {
				break ILOOP
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	// 写入输出流
	w.Write([]byte(Message))
	w.Flush()
	// 打印发送的消息
	log.Printf("Send: %s", Message)

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

func main() {
	// 监听端口号
	port := 3333
	// 启动服务，Socket套接字网络传输
	SocketServer(port)
}
