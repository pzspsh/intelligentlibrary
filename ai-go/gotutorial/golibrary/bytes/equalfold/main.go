/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:42:51
*/
package main

import (
	"bytes"
	"fmt"
)

/*
EqualFold报告解释为UTF-8字符串的s和t在简单的Unicode大小写折叠下是否相等，这是一种更通用的大小写不敏感形式。
*/
func main() {
	fmt.Println(bytes.EqualFold([]byte("Go"), []byte("go"))) // true
	fmt.Println(bytes.EqualFold([]byte{}, []byte{}))         // true
	fmt.Println(bytes.EqualFold([]byte{'A'}, []byte{'a'}))   // true
	fmt.Println(bytes.EqualFold([]byte{'B'}, []byte{'a'}))   // false
	fmt.Println(bytes.EqualFold([]byte{}, nil))              // true
}
