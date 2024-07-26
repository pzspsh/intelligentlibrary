/*
@File   : main.go
@Author : pan
@Time   : 2024-07-26 17:25:46
*/
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	proxyConn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		fmt.Println("Error connecting to proxy: ", err.Error())
		return
	}
	defer proxyConn.Close()

	fmt.Println("Connected to google.com")

	go func() {
		defer conn.Close()
		defer proxyConn.Close()

		io.Copy(conn, proxyConn)
	}()

	go func() {
		defer conn.Close()
		defer proxyConn.Close()

		io.Copy(proxyConn, conn)
	}()
}
