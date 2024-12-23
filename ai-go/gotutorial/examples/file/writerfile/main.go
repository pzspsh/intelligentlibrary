/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:40:07
*/
package main

import (
	"log"
	"os"
)

func main() {
	// 可写方式打开文件
	file, err := os.OpenFile(
		"test.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写字节到文件中
	byteSlice := []byte("Bytes!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)
}
