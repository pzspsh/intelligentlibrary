/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:41:23
*/
package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
)

func main() {
	// 一个缓冲区存储压缩的内容
	buf := bytes.NewBuffer(nil)

	// 创建一个flate.Writer，压缩级别最好
	flateWriter, err := flate.NewWriterDict(buf, flate.BestCompression, []byte("key"))
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWriter.Close()
	// 写入待压缩内容
	flateWriter.Write([]byte("compress/flate\n"))
	flateWriter.Flush()
	fmt.Println(buf)
}
