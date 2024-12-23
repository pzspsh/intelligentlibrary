/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:38:49
*/
package main

import (
	"bytes"
	"compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 一个缓存区压缩的内容
	buf := bytes.NewBuffer(nil)

	// 创建一个flate.Writer
	flateWrite, err := flate.NewWriter(buf, flate.BestCompression)
	if err != nil {
		log.Fatalln(err)
	}
	defer flateWrite.Close()
	// 写入待压缩内容
	flateWrite.Write([]byte("compress/flate\n"))
	flateWrite.Flush()
	fmt.Printf("压缩后的内容：%s\n", buf)

	// 解压刚压缩的内容
	flateReader := flate.NewReader(buf)
	defer flateWrite.Close()
	// 输出
	fmt.Print("解压后的内容：")
	io.Copy(os.Stdout, flateReader)
}
