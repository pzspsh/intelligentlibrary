/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:02:58
*/
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		//handle
		fmt.Println(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}
		// defer conn.Close() // 这个无限循环中的deferred永远不会运行
		for {
			//处理
			reader := bufio.NewReader(conn)
			var buf [1024]byte
			read, err := reader.Read(buf[:])
			if err != nil {
				//handle
				fmt.Println(err)
			}
			recv := string(buf[:read])
			fmt.Println("GET MESSAGE: ", recv)
			conn.Write([]byte(recv))
		}
	}
}
