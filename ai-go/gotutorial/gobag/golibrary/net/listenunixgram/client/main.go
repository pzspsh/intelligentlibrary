/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 20:08:54
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	raddr, err := net.ResolveUnixAddr("unixgram", "/tmp/unix_test_sock")
	checkError(err)

	laddr, err := net.ResolveUnixAddr("unixgram", "/tmp/unix_test_sock_cli")
	checkError(err)

	conn, err := net.DialUnix("unixgram", laddr, raddr)
	checkError(err)

	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}

	n, err := conn.Write([]byte("Hello world"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("send msg n:%d\n", n)

	var msg [20]byte
	conn.Read(msg[0:])

	fmt.Println("msg is", string(msg[0:10]))
}
