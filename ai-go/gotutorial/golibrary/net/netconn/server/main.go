/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:21:37
*/
package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close()
	for {
		//创建一个新的切片
		buf := make([]byte, 1024)
		//等待客户端通过conn发送信息
		//如果客户端没有write[发送], 那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Printf("客户端退出 err=%v\n", err)
			return
		}
		//3. 显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n]))
	}
}
func main() {
	fmt.Println("服务器开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()
	//循环等待客户端来连接
	for {
		fmt.Println("等待客户端来连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}
		//准备一个协程，处理客户端
		go process(conn)
	}
	//fmt.Printf("listen=%v\n", listen)
}
