/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 00:24:05
*/
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%t\n", unicode.IsLetter('A'))
	fmt.Printf("%t\n", unicode.IsLetter('7'))
}