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
	var servadd *net.UnixAddr

	servadd, _ = net.ResolveUnixAddr("unix", "D:\\server_sock")

	ulistener, _ := net.ListenUnix("unix", servadd)

	defer ulistener.Close()

	for {
		conn, err := ulistener.AcceptUnix()
		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + conn.RemoteAddr().String())
		go unixrw(conn)
	}

}

func unixrw(conn *net.UnixConn) {
	addrStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + addrStr)
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
