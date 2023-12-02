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
	str := "世"
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))
}
