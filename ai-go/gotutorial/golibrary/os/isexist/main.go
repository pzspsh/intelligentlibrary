/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 15:58:26
*/
package main

import (
	"fmt"
	"os"
)

/* 根据错误，判断 文件或目录是否存在 */
func main() {
	if _, err := os.Open("filepath/test.js"); err != nil {
		// false 不存在   true 存在
		emptyErr := os.IsExist(err) // 常用于这些os.OpenFile、os.Create、os.Mkdir函数
		fmt.Println(emptyErr, "\n", err)
	}
}
