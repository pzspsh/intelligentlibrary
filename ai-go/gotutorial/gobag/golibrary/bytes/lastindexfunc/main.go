/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 14:48:41
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
lastindexfunc c将s解释为utf -8编码的代码点序列。它返回s中最后一个
满足f(c)的Unicode码点的字节索引，如果不满足则返回-1。
*/
func main() {
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsLetter))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsPunct))
	fmt.Println(bytes.LastIndexFunc([]byte("go gopher!"), unicode.IsNumber))
}
