/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 15:57:36
*/
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
