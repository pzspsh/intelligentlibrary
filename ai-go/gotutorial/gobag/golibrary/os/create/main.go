/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:06:22
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("文件创建成功")
}
