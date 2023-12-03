/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:47:13
*/
package main

import (
	"log"
	"net"
)

func main() {
	//尝试连接百度服务器
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		log.Fatal("连接失败!", err)
	}
	defer conn.Close()
	log.Println("连接成功!")
	//发送HTTP形式的内容
	conn.Write([]byte("GET / HTTP/1.1\r\nHost: www.baidu.com\r\nUser-Agent: curl/7.55.1\r\nAccept: ＊/＊\r\n\r\n"))

	var buf = make([]byte, 1024)
	conn.Read(buf)
	log.Println(string(buf))
}
