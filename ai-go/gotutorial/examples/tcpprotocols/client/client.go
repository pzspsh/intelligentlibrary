/*
@File   : client.go
@Author : pan
@Time   : 2024-07-11 13:12:11
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 1. 建立连接
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Failed to Dial")
		return
	}
	// 延迟关闭，释放资源
	defer conn.Close()

	// 2. 与服务端通信
	buf := make([]byte, 256)
	for {
		// 2.1 从控制台读取输入
		readBytesCount, _ := os.Stdin.Read(buf)
		// 2.2 写到网络(即：发送请求)
		conn.Write(buf[:readBytesCount])
		// 2.3 读网络(即：获取响应)
		readBytesCount, _ = conn.Read(buf)
		// 2.4 输出到控制台
		// 提示：buf[:n]的效果为：读取buf[总长度-n]至buf[n]处的字节
		os.Stdout.Write(buf[:readBytesCount])
	}
}
