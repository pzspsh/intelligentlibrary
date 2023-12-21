/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:04:30
*/
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
