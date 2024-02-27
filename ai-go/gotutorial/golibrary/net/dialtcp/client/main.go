/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 19:16:30
*/
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9999,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	//	发送数据请求
	writer := bufio.NewWriter(conn)
	req := "Hello,server!\n"
	_, err = writer.WriteString(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	writer.Flush()

	//	接受响应数据
	reader := bufio.NewReader(conn)
	resp, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("接收到达额数据是", resp)
}
