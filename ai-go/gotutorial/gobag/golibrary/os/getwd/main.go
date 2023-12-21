/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:01:36
*/
package main

import (
	"fmt"
	"os"
)

/* 获取当前工作目录 */
func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
}
