/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(bytes.Map(func(r rune) rune {
		return r + 1 // 将每一个字符都+1
	}, []byte("abc")))) // bcd
}
