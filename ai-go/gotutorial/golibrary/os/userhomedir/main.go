/*
@File   : main.go
@Author : pan
@Time   : 2024-05-30 16:09:07
*/
package main

import (
	"fmt"
	"os"
)

// 获取系统用户的根目录

func main() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User's home directory:", userHomeDir)
}
