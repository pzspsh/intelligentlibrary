/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:40:46
*/
package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"log"
)

func main() {
	// 一个缓冲区压缩的内容
	buf := bytes.NewBuffer(nil)

	// 创建一个flate.Writer，压缩级别最好
	flateWrite, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWrite.Close()
	// 写入待压缩内容
	flateWrite.Write([]byte("compress/flate\n"))
	flateWrite.Flush()
	fmt.Println(buf)
}
