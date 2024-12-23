/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:04:05
*/
package main

import (
	"bytes"
	"fmt"
)

/*
ToValidUTF8将s视为UTF-8编码的字节，并返回一个副本，其中每运行一次表示
无效UTF-8的字节都替换为replacement中的字节，replace中的字节可能为空。
*/
func main() {
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("abc"), []byte("\uFFFD")))           // abc
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("a\xffb\xC0\xAFc\xff"), []byte(""))) // abc
	fmt.Printf("%s\n", bytes.ToValidUTF8([]byte("\xed\xa0\x80"), []byte("abc")))     // abc
}
