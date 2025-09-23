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
