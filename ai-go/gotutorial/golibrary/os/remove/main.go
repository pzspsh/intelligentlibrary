package main

import (
	"fmt"
	"os"
)

/* 移除文件或目录(单一文件) */
func main() {
	if err := os.Remove("test"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
