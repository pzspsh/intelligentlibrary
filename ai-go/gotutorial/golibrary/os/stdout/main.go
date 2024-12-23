/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 11:00:42
*/
package main

import (
	"fmt"
	"log"
	"os"
)

type abcReader struct {
}

func (r *abcReader) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

type abcWriter struct {
}

func (r *abcWriter) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func main() {
	var reader abcReader
	var writer abcWriter
	fmt.Printf("please input a string.\n")
	//创建缓冲区，保存输入输出
	input := make([]byte, 4096)
	//使用reader 读取输入
	n, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}

	fmt.Printf("read %d bytes from stdin\n", n)
	//使用writer写入输出
	n2, err := writer.Write(input[:n])
	if err != nil {
		log.Fatalln("Unable to write data")
	}
	fmt.Printf("write %d bytes to stout\n", n2)
	/**
	please input a string.
	i love you
	read 12 bytes from stdin
	i love you
	write 12 bytes to stout
	*/
}
