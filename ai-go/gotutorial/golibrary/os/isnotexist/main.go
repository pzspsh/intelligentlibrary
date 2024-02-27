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
	if _, err := os.Open("filepath/test.js"); err != nil {
		// false 不存在   true 存在
		emptyErr := os.IsNotExist(err)
		fmt.Println(emptyErr, "\n", err)
	}
}
