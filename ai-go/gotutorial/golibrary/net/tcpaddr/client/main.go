/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:46:31
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	clinet()
}

func clinet() {
	//发送地址
	socket, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})
	if err != nil {
		fmt.Printf("dial%s\n", err)
		return
	}
	//关闭连接
	defer socket.Close()

	for {
		sendData := []byte("hello server")
		//写入连接
		_, err = socket.Write(sendData)
		if err != nil {
			fmt.Printf("send%s\n", err)
			return
		}
		Data := make([]byte, 1024)
		//读出数据
		_, err = socket.Read(Data)
		if err != nil {
			fmt.Printf("read%s\n", err)
			return
		}
		fmt.Printf("%s\n", Data)
	}
}
