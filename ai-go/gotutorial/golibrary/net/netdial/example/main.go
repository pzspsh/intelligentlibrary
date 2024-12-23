/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:46:35
*/
package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		log.Fatal("连接失败！", err)
	}
	defer conn.Close()
	log.Println("连接成功！")
	//发送数据
	conn.Write([]byte("test\n"))
	//接收数据
	var buf = make([]byte, 10)
	conn.Read(buf)
	log.Println(buf)
}
