/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:45
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func SendReceiveProccess(conn net.Conn) {
	// 3.与客户端进行消息通信(循环读取客户端法过来的信息)
	defer conn.Close() // 关闭当前链接通信对象
	reader := bufio.NewReader(conn)
	var msg [1024]byte // 每次读取1024B
	for {
		n, err := reader.Read(msg[:]) // 注意读取的类型是byte的slice
		// 末尾标识
		if err == io.EOF {
			fmt.Printf("Close conn %v\n", conn.RemoteAddr())
			break
		}
		// 异常时break
		if err != nil {
			fmt.Printf("Read from Client conn failed, Close conn %v\n", conn.RemoteAddr())
			break
		}
		fmt.Println(conn.RemoteAddr(), "->", string(msg[:n]))

		// 将客户端发送的信息又转发给客户端(返回写入的字节数)
		_, err = conn.Write([]byte(msg[:n]))
		if err != nil {
			fmt.Printf("Send failed, Close Client conn %v\n", conn.RemoteAddr())
			break
		}
	}
}

func main() {
	// 1.设置监听端口
	address := "10.20.172.108:22022"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Start Tcp Server on %v Failed!\nerr：%v\n", address, err)
		return
	} else {
		fmt.Printf("Server Listen %v Start......\n", address)
	}
	defer listener.Close() // 关闭服务端监听

	// 2.等待客户端建立连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept failed,err: %v\n", err)
			return
		}
		// 不同的客户端利用Goroutine分配不同的线程进行响应。
		go SendReceiveProccess(conn)
	}
}
