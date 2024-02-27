/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:09:14
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("文件打开成功")
}
