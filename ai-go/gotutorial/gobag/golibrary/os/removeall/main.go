/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:05:09
*/
package main

import (
	"fmt"
	"os"
)

/* 递归删除文件或目录 */
func main() {
	if err := os.RemoveAll("test01"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
