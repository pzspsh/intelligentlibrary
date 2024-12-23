/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:48:39
*/
package main

import (
	"encoding/base32"
	"fmt"
	"strings"
)

func main() {
	src := "ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA===="
	reader := strings.NewReader(src)
	dst := ""
	decoder := base32.NewDecoder(base32.StdEncoding, reader)
	// 使用一个很小的输出buffer，测试流式解码
	buf := make([]byte, 2)
	for {
		n, err := decoder.Read(buf)
		if err != nil || n == 0 {
			break
		}
		dst += string(buf[:n])
		fmt.Println(dst)
	}
}
