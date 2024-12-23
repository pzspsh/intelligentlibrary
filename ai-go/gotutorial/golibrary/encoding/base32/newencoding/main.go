/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:49:52
*/
package main

import (
	"encoding/base32"
	"fmt"
)

func main() {
	encodeTest := "--------------------------------"
	enc := base32.NewEncoding(encodeTest)
	src := "this is a test string."
	dst := enc.EncodeToString([]byte(src))

	// 最后不足8字节的会用"="补全
	fmt.Println(dst)
	fmt.Println(len(dst))
}
