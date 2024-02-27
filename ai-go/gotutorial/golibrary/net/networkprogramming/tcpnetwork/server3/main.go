/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:58:45
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	checkErr(err)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		checkErr(err)
		go task(conn)
	}
}
func task(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect sucessful")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		checkErr(err)
		if string(buf[:n]) == "exit" {
			conn.Write([]byte("拜拜~~~~"))
			conn.Close()
			fmt.Println(addr, " Disconnect sucessful")
			return
		}
		fmt.Printf("read buf = %s\n", string(buf[:n]))
		response := "Hello " + string(buf[:n])
		conn.Write([]byte(response))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
