/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:23:49
*/
package main

import (
	"fmt"
	"io"
	"net"
)

/*
func main() {
	//建立socket，监听端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)
	}
}

// 处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
}

func Log(v ...interface{}) {
	log.Println(v...)
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
*/

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
