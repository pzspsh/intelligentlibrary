/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:01:07
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
totitlspecial将s视为utf -8编码的字节，并返回一个副本，
其中包含映射到其标题大小写的所有Unicode字母，并优先考虑特殊的大小写规则。
*/
func main() {
	str := []byte("ahoj vývojári golang")
	totitle := bytes.ToTitleSpecial(unicode.AzeriCase, str)
	fmt.Println("Original : " + string(str))    // Original : ahoj vývojári golang
	fmt.Println("ToTitle : " + string(totitle)) // ToTitle : AHOJ VÝVOJÁRİ GOLANG
}
