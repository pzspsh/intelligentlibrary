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
	fmt.Println(string(bytes.ToLowerSpecial(unicode.SpecialCase{}, []byte("aAA")))) // aaa
	str := []byte("AHOJ VÝVOJÁRİ GOLANG")
	totitle := bytes.ToLowerSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))    // Original : AHOJ VÝVOJÁRİ GOLANG
	fmt.Println("ToLower : " + string(totitle)) // ToLower : ahoj vývojári golang
}
