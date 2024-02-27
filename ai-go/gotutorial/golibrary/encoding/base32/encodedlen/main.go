/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:51:34
*/
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	src := []byte("this is a test string.")
	maxLen := base32.StdEncoding.EncodedLen(len(src))
	dst := make([]byte, maxLen)
	base32.StdEncoding.Encode(dst, src)
	fmt.Println("编码后的数据为:", string(dst))
}
