/*
@File   : mian.go
@Author : pan
@Time   : 2023-12-03 16:58:26
*/
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	//建立socket端口监听
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()
	Log("Waiting for clients ...")

	//等待客户端访问
	for {
		conn, err := netListen.Accept() //监听接收
		if err != nil {
			continue //如果发生错误，继续下一个循环。
		}
		Log(conn.RemoteAddr().String(), "tcp connect success") //tcp连接成功
		go handleConnection(conn)
	}
}

// 处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048) //建立一个slice
	for {
		n, err := conn.Read(buffer) //读取客户端传来的内容
		if err != nil {
			Log(conn.RemoteAddr().String(), "connection error: ", err)
			return //当远程客户端连接发生错误（断开）后，终止此协程。
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

		//返回给客户端的信息
		strTemp := "CofoxServer got msg \"" + string(buffer[:n]) + "\" at " + time.Now().String()
		conn.Write([]byte(strTemp))
	}
}

// 日志处理
func Log(v ...interface{}) {
	log.Println(v...)
}

// 错误处理
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}
}
