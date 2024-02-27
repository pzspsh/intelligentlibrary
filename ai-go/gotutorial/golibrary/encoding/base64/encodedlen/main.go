/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:59:44
*/
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	src := []byte("this is a test string.")
	maxLen := base64.StdEncoding.EncodedLen(len(src))
	fmt.Println("编码后的数据长度为:", maxLen)
}
