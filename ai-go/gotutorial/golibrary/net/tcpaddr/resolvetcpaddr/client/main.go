/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:47:55
*/
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//1.创建链接远程链接服务器，得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}
	i := 1
	for {
		//2.调用链接Write写数据
		_, err := conn.Write([]byte(fmt.Sprintf("%s:%d", "Hello Server", i)))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err")
			return
		}
		fmt.Printf("Server call back:%s,cnt = %d\n", buf, cnt)
		i++
		time.Sleep(1 * time.Second)
	}
}
