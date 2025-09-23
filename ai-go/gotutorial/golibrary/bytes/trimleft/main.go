package main

import (
	"bytes"
	"fmt"
)

/*
通过切掉cutset中包含的所有前导utf -8编码的代码点，TrimLeft返回s的子切片。
*/
func main() {
	fmt.Print(string(bytes.TrimLeft([]byte("453gopher8257"), "0123456789"))) // gopher8257
	fmt.Println(string(bytes.TrimLeft([]byte("hi hi go"), "hi")))            //  hi go
}
