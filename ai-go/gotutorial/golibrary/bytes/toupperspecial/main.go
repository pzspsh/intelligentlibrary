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

/*
ToUpperSpecial将s视为utf -8编码的字节，并返回一个副本，其中所有Unicode字母映射到它们的大写，
优先考虑特殊的大小写规则。
*/
func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToUpperSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))                                        // Original : ahoj vývojári golang
	fmt.Println("ToUpper : " + string(totitle))                                     // ToUpper : AHOJ VÝVOJÁRİ GOLANG
	fmt.Println(string(bytes.ToUpperSpecial(unicode.SpecialCase{}, []byte("Aaa")))) // AAA
}
