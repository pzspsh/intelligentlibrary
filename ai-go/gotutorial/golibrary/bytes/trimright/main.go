package main

import (
	"bytes"
	"fmt"
)

/*
TrimRight通过切掉cutset中包含的所有尾随utf -8编码的代码点，返回s的子切片。
*/
func main() {
	fmt.Print(string(bytes.TrimRight([]byte("453gopher8257"), "0123456789"))) // 453gopher
}
