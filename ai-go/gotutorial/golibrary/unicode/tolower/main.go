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
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG))
}
