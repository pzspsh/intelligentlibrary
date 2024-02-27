/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 15:07:37
*/
package main

import (
	"fmt"
	"io"
	"net"
)

/*
net.Buffers.Read()将实现io.Reader。请阅读文档Read()的含义以及它的工作原理。

简而言之，Read()方法用于将数据读取到您传递给它的片中，读取的数据将被写入该片。

要从net.Buffers中读取数据，首先要将一些数据“放入”其中。net.Buffers是一片切片：

type Buffers [][]byte
*/
func main() {
	var bufs net.Buffers
	bufs = append(bufs, []byte("Hello"))
	bufs = append(bufs, []byte(" World"))
	data := make([]byte, 100)
	n, err := bufs.Read(data)
	if err != nil {
		if err == io.EOF {
			// This just means everything from bufs was read,
			// and data isn't filled completely
		} else {
			fmt.Println("bufs.Read error:", err)
			return
		}
	}
	fmt.Printf("Read %d bytes\n", n)
	fmt.Printf("Data: %s\n", data[:n])
}
