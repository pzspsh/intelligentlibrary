/*
@File   : main.go
@Author : pan
@Time   : 2024-07-11 13:12:02
*/
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// 1. 绑定ip和端口，设置监听
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to Listen", err)
	}
	// 延迟关闭，释放资源
	defer listener.Close()

	// 2. 循环等待新连接
	for {
		// 从连接列表获取新连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept", err)
		}
		// 3. 与新连接通信(为了不同步阻塞，这里开启异步协程进行函数调用)
		go handle_conn(conn)
	}
}

/**
 * 处理连接
 */
func handle_conn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New connection ", conn.RemoteAddr())
	// 通信
	buf := make([]byte, 256)
	for {
		// 从网络中读
		readBytesCount, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read", err)
			break
		}
		// 提示：buf[:n]的效果为：读取buf[总长度-n]至buf[n]处的字节
		fmt.Println("服务端收到数据：\t", string(buf[:readBytesCount]))

		// 写回网络 -- 收到什么就写回什么，即：回射服务器
		writeByteCount, err := conn.Write(buf[:readBytesCount])
		if err != nil {
			fmt.Println("Failed to write", err)
			break
		}
		fmt.Printf("write success %d bytes\n", writeByteCount)
	}
}
