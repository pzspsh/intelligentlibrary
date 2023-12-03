/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 18:46:31
*/
package main

import (
	"fmt"
	"net"
)

func main() {
	server()
}

func server() {
	//监听地址
	listen, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8080,
	})

	if err != nil {
		fmt.Printf("错误：%s\n", err)
		return
	}
	//关闭
	defer listen.Close()

loop:
	// 等待连接
	con, err := listen.AcceptTCP()

	if err != nil {
		fmt.Printf("accepttcp  %s\n", err)

	} else {
		// 没有错误就执行收发操作，这里用协程 方便多个连接进来
		go Line(con)
	}
	// 跳转到继续等待
	goto loop

}
func Line(con *net.TCPConn) {

	for {

		var data [1024]byte
		sendData := []byte("lixing")
		// 读连接里面的数据
		_, err := con.Read(data[:])

		if err != nil {
			fmt.Printf("read   %s\n", err)
			con.Close()
			return
		}

		fmt.Printf("%s\n", data)
		//向连接写入数据
		_, err = con.Write(sendData)

		if err != nil {
			fmt.Printf("write%s\n", err)
			con.Close()
			return
		}

	}
}
