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
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})

	fmt.Println(utf8.ValidString(valid))
	fmt.Println(utf8.ValidString(invalid))
}
