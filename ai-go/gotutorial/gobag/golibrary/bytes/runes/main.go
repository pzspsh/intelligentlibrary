/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:43:14
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println([]byte("你好world"))              // [228 189 160 229 165 189 119 111 114 108 100]
	fmt.Println(bytes.Runes([]byte("你好world"))) // [20320 22909 119 111 114 108 100]
}
