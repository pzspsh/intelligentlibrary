/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:16:42
*/
package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleClient(conn *net.TCPConn) {
	defer conn.Close()

	//	读取客户端请求
	reader := bufio.NewReader(conn)
	req, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("接收到的请求是：", req)

	//发送响应数据

	writer := bufio.NewWriter(conn)
	resp := "Hello,client!\n"
	_, err = writer.WriteString(resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Flush()

}

func main() {

	tcpListenner, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9999,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tcpListenner.Close()

	fmt.Println("监听端口:", 9999)

	for {
		conn, err := tcpListenner.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("新的客户端连接...")
		go handleClient(conn)

	}
}
