/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:27:58
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	h := sha256.New()
	h.Write([]byte(s))
	// 计算SHA256哈希
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
