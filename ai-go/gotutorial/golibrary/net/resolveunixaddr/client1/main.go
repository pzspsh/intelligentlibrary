/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 20:01:57
*/
package main

import (
	"fmt"
	"net"
	"time"
)

var quitSig chan bool

func main() {
	var serverAddr *net.UnixAddr                                   //定义UNIX domain socket地址
	serverAddr, _ = net.ResolveUnixAddr("unix", "D:\\server_sock") //服务地址
	conn, _ := net.DialUnix("unix", nil, serverAddr)               //"unix"表示SOCK_STREAM方式
	defer conn.Close()
	fmt.Println("connected!")

	go ClientSend(conn) //开启发送线程
	b := []byte("cleint send example\n")
	conn.Write(b)

	<-quitSig
}

func ClientSend(conn *net.UnixConn) {
	for {
		time.Sleep(time.Second)
		//创建消息缓冲区
		buffer := make([]byte, 1024*1024*8)
		buffer[0] = '1'
		t := time.Now()
		conn.Write(buffer) //发送8M数据
		elapsed := time.Since(t)
		fmt.Println("app elapsed:", elapsed)
	}
}
