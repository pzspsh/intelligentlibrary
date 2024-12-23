/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:45:29
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
indexfunc c将s解释为utf-8编码的代码点序列。它返回s中第一个满足f(c)的Unicode
码点的字节索引，如果不满足则返回-1。
*/
func main() {
	fmt.Println(bytes.IndexFunc([]byte("hi go"), func(r rune) bool {
		return r == 'g'
	})) // 3
	fmt.Println("##############################")
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f))    // 7
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f)) // -1
}
