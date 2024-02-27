/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 20:01:57
*/
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

var quitSemaphore chan bool

func main() {
	var addr *net.UnixAddr
	addr, _ = net.ResolveUnixAddr("unix", "/tmp/unix_ss")

	conn, _ := net.DialUnix("unix", nil, addr)
	defer conn.Close()
	fmt.Println("connected!")

	go onMessageRecived(conn)

	b := []byte("time\n")
	conn.Write([]byte("client: "))
	conn.Write(b)

	<-quitSemaphore
}

func onMessageRecived(conn *net.UnixConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil {
			quitSemaphore <- true
			break
		}
		time.Sleep(time.Second)
		b := []byte(msg)
		conn.Write([]byte("client: "))
		conn.Write(b)
	}
}
