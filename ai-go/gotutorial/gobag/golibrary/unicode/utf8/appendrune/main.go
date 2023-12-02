/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 00:32:09
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	buf1 := utf8.AppendRune(nil, 0x10000)
	buf2 := utf8.AppendRune([]byte("init"), 0x10000)
	fmt.Println(string(buf1))
	fmt.Println(string(buf2))
}
