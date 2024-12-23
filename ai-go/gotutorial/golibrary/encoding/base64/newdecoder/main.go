/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:57:04
*/
package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	// base 64 数据
	src := "dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg=="
	reader := strings.NewReader(src)
	decoder := base64.NewDecoder(base64.StdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 2)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		if n == 0 || err != nil {
			break
		}
		dst += string(buf[:n])
	}
	fmt.Println("解码后的数据为:", dst)
}
