package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("连接失败，错误：", err)
		return
	}
	defer conn.Close()
	for {
		var name string
		_, _ = fmt.Scanln(&name)
		conn.Write([]byte(name))

		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)
		fmt.Println(string(buf[:n]))
	}
}
