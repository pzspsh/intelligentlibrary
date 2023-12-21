/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:00:28
*/
package main

import (
	"fmt"
	"os"
)

/* 根据错误，判断是否为权限错误 */
func main() {
	file, _ := os.Open("filepath/test.js")
	_, err := file.WriteString("// new info")
	if err != nil {
		fmt.Println(os.IsPermission(err))
	}
	defer file.Close()
}
