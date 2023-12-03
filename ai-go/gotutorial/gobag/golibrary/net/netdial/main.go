/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:43:35
*/
package main

import (
	"log"
	"net"
)

func main() {
	// 尝试连接本地1234端口
	// conn, err := net.Dial("tcp", ":1234")
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal("连接失败！", err)
	}
	defer conn.Close()
	log.Println("连接成功！")
}
