/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:56:31
*/
package main

import (
	"bytes"
	"fmt"
)

/*
Fields将s解释为utf -8编码的码点序列。按照unicode的定义，它围绕一个或多个连续空格字
符的每个实例拆分切片。IsSpace，返回s的子切片的切片，如果s只包含空白则返回空切片。
*/
func main() {
	s := bytes.Fields([]byte(" hi 你啊,    is not good, my boy"))
	for _, v := range s {
		fmt.Print(string(v) + "|") // hi|你啊,|is|not|good,|my|boy|
	}
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("  foo bar  baz   ")))
}
