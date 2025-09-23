package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("path/file.txt"); err != nil {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件存在")
	}
}
