/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:45:33
*/
package main

import (
	"log"
	"os"
)

func main() {
	// 打开文件，只读
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 从文件中读取len(b)字节的文件。
	// 返回0字节意味着读取到文件尾了
	// 读取到文件会返回io.EOF的error
	byteSlice := make([]byte, 16)
	bytesRead, err := file.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}
