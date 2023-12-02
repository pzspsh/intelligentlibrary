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
	runes := []rune{
		// Less than 0, out of range.
		-1,
		// Greater than 0x10FFFF, out of range.
		0x110000,
		// The Unicode replacement character.
		utf8.RuneError,
	}
	for i, c := range runes {
		buf := make([]byte, 3)
		size := utf8.EncodeRune(buf, c)
		fmt.Printf("%d: %d %[2]s %d\n", i, buf, size)
	}
}
