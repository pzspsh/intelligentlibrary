/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:48:38
*/
package main

import (
	"log"
	"os"
)

func main() {
	// 读取文件到byte slice中
	data, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}
