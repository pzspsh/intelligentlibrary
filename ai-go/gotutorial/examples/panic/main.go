/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:10:08
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	panic("a problem")
	_, err := os.Create("/tmp/file")
	if err != nil {
		// 当遇到error不知道怎么处理时，可以包装成panic
		fmt.Println(err)
		panic(err)
	}
}
