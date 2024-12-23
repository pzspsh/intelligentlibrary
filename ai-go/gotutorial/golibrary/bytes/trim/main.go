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

/*
Trim通过切掉cutset中包含的所有开头和结尾的utf -8编码的代码点，返回s的子切片。
*/
func main() {
	fmt.Printf("[%q]", bytes.Trim([]byte(" !!! Achtung! Achtung! !!! "), "! ")) // ["Achtung! Achtung"]
	fmt.Println(string(bytes.Trim([]byte("hello my"), "my")))                   // hello
}
