/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:15:18
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("filepath/hello.go", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString("//直接写入的字符串数据")
}
