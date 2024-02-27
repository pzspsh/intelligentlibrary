/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:39:03
*/
package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//获取客户端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect successful")
	buf := make([]byte, 2048)
	defer conn.Close()
	for {

		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.read=", err)
			return
		}
		len -= 2
		fmt.Println("buf=", string(buf[:len]))
		//fmt.Println(len, buf)
		if string(buf[:len]) == "exit" {
			exit := []byte(addr)
			exit = append(exit, "is exit"...)
			conn.Write(exit)
			fmt.Println(addr, "exit")
			return
		}
		//处理操作，小写变大写
		conn.Write([]byte(strings.ToUpper(string(buf[:len]))))
	}
}

func main() {
	//监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	//阻塞等待用户连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//接收用户请求
		go HandleConn(conn)
	}

}
