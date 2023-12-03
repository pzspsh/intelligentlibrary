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

func main() {
	var addr *net.UnixAddr

	addr, _ = net.ResolveUnixAddr("unix", "/tmp/unix_ss")

	unixListener, _ := net.ListenUnix("unix", addr)

	defer unixListener.Close()

	for {
		conn, err := unixListener.AcceptUnix()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + conn.RemoteAddr().String())
		go unixPipe(conn)
	}

}

func unixPipe(conn *net.UnixConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		//读取客户端内容
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		//将当前时间写回给客户端
		conn.Write([]byte("service : "))
		conn.Write(b)
	}
}
