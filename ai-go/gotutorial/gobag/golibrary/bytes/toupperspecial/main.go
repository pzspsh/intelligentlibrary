/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:50:26
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(string(bytes.ToUpperSpecial(unicode.SpecialCase{}, []byte("Aaa")))) // AAA
}
