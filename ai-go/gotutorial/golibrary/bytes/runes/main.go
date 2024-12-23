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

/*
Runes将s解释为utf -8编码的代码点序列。它返回一段符文(Unicode码点)，相当于s。
*/
func main() {
	fmt.Println([]byte("你好world"))              // [228 189 160 229 165 189 119 111 114 108 100]
	fmt.Println(bytes.Runes([]byte("你好world"))) // [20320 22909 119 111 114 108 100]

	rs := bytes.Runes([]byte("go gopher"))
	for _, r := range rs {
		fmt.Printf("%#U\n", r)
	}
}
