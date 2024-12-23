/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:44:55
*/
package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io"
	"os"
)

func main() {
	// 一个缓冲区存储压缩的内容
	buf := bytes.NewBuffer(nil)

	w := lzw.NewWriter(buf, lzw.LSB, 8)
	// 写入待压缩内容
	w.Write([]byte("compress/flate\n"))
	w.Close()
	fmt.Println(buf)

	// 解压
	r := lzw.NewReader(buf, lzw.LSB, 8)
	defer r.Close()
	io.Copy(os.Stdout, r)
}
