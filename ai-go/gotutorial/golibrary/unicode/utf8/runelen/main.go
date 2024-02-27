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
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('界'))
}
