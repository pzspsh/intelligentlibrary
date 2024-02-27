/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:40:00
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//接收服务器信息
	go func() {
		buf := make([]byte, 2045)
		for {
			len, err2 := conn.Read(buf)
			if err2 != nil {
				fmt.Println("conn.Read err=", err)
				break
			}
			fmt.Println(string(buf[:len]))
		}
	}()
	//发送数据
	for {
		buf := make([]byte, 1024)
		fmt.Scan(&buf)
		if err != nil {
			fmt.Println("os.Stdin err=", err)
		}
		conn.Write(buf[0:3])
	}

}
