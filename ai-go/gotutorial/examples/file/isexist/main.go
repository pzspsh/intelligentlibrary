/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:32:54
*/
package main

import (
	"log"
	"os"
)

var (
	fileInfo *os.FileInfo
	err      error
)

func main() {
	// 文件不存在则返回error
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:")
	log.Println(fileInfo)
}
