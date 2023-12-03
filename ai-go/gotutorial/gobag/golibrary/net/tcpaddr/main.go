/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:44:26
*/
package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// tcp echo server，向客户端以固定周期发送时间戳字符串
func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second) // 五秒的心跳间隔
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}

// tcp客户端：读取echo server发送的时间戳
func client() {
	time.Sleep(3 * time.Second)
	/*if len(os.Args) != 2 {
	  log.Fatalf("Usage: %s host:port", os.Args[0])
	}*/
	//service := os.Args[1]
	service := "127.0.0.1:8000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err, n)
		}
		if err != nil {
			log.Fatal(err)
			fmt.Println("client write message error")
		} else {
			fmt.Println("read from server, message:", string(buf[:]))
		}

		n, err = conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
		if err != nil {
			log.Fatal(err)
			fmt.Println("client write message error")
		}
		time.Sleep(10 * time.Second)
		//log.Fatal(n)
		fmt.Println(n)
	}
}

func main() {

	go client()

	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"), // 把字符串IP地址转换为net.IP类型
		Port: 8000,
	}
	listener, err := net.ListenTCP("tcp4", &address) // 创建TCP4服务器端监听器
	if err != nil {
		log.Fatal(err) // Println + os.Exit(1)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err) // 错误直接退出
		}
		fmt.Println("remote address:", conn.RemoteAddr())
		go echo(conn)
	}
}
