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

func recvUnixMsg(conn *net.UnixConn) {
	var buf [20]byte

	n, raddr, err := conn.ReadFromUnix(buf[0:])
	fmt.Println(raddr)
	if err != nil {
		return
	}

	fmt.Println("msg is ", string(buf[0:n]))
	_, err = conn.WriteToUnix([]byte("nice to see u"), raddr)
	checkError(err)
}

func main() {
	laddr, err := net.ResolveUnixAddr("unixgram", "/tmp/unix_gram_sock")
	checkError(err)

	conn, err := net.ListenUnixgram("unixgram", laddr)
	checkError(err)

	recvUnixMsg(conn)
}
