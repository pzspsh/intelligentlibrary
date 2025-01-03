/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:57:53
*/
package main

import (
	"encoding/base64"
	"fmt"
)

// 我们的编码表，64 字节
const encodeTest = "----------------------------------------------------------------"

func main() {
	enc := base64.NewEncoding(encodeTest)
	src := "this is a test string."
	dst := enc.EncodeToString([]byte(src))

	// 最后不足4字节的会用"="补全
	fmt.Println(dst)
	fmt.Println(len(dst)%4 == 0)
}
