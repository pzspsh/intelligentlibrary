/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:05:47
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// 重命名
	err := os.Rename("test.txt", "test01.js")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Mkdir("test", os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	// 移动
	err = os.Rename("test01.js", "test/text01.txt")
	if err != nil {
		fmt.Println(err)
	}
}
