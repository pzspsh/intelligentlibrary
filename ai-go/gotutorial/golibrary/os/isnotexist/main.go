/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 15:59:07
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	path := "filepath/test.js"
	if _, err := os.Open(path); err != nil {
		// false 不存在   true 存在
		if _, emptyErr := os.Stat(path); os.IsNotExist(emptyErr) {
			fmt.Println("不存在")
		} else {
			fmt.Println("存在")
		}
		emptyErr := os.IsNotExist(err) // 用户os.Stat函数
		fmt.Println(emptyErr, "\n", err)
	}
}
