/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 14:58:40
*/
package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	src := []byte("dGhpcyBpcyBhIHRlc3Qgc3RyaW5nLg==")
	maxLen := base64.StdEncoding.DecodedLen(len(src))
	dst := make([]byte, maxLen)
	n, err := base64.StdEncoding.Decode(dst, src)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解码后的数据为:", string(dst[:n]))
}
