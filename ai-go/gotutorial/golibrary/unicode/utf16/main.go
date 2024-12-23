/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:04:43
*/
package main

import (
	"fmt"
	"unicode/utf16"
)

func main() {
	words := []rune{'𝓐', '𝓑'}
	u16 := utf16.Encode(words)
	fmt.Println(u16)
	fmt.Println(utf16.Decode(u16))
	r1, r2 := utf16.EncodeRune('𝓐')
	fmt.Println(r1, r2)
	fmt.Println(utf16.DecodeRune(r1, r2))
	fmt.Println(utf16.IsSurrogate(r1))
	fmt.Println(utf16.IsSurrogate(r2))
	fmt.Println(utf16.IsSurrogate(1234))
}
