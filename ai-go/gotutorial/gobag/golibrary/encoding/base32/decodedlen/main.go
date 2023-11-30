/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:50:25
*/
package main

import (
	"encoding/base32"
	"fmt"
	"log"
)

func main() {
	src := []byte("ORUGS4ZANFZSAYJAORSXG5BAON2HE2LOM4XA====")
	// 解码后数据的最长长度
	maxLen := base32.StdEncoding.DecodedLen(len(src))
	// 解码后的缓存区
	dst := make([]byte, maxLen)
	// base32 解码
	n, err := base32.StdEncoding.Decode(dst, src)
	// 打印解码的数据
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解码后的数据:", string(dst[:n]))
}
