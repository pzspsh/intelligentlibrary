/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 00:08:18
*/
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:10051")

	if err != nil {
		panic(err)
	}

	for {
		conn, _ := ln.Accept() //The loop will be held here
		fmt.Println("get connect")
		go handleread(conn)

	}
}

func handleread(conn net.Conn) {
	defer conn.Close()

	var tatalBuffer []byte
	var all int
	for {
		buffer := make([]byte, 2)
		n, err := conn.Read(buffer)
		if err == io.EOF {
			fmt.Println(err, n)
			break
		}

		tatalBuffer = append(tatalBuffer, buffer...)
		all += n

		fmt.Println(string(buffer), n, string(tatalBuffer[:all]), all)
	}
}
