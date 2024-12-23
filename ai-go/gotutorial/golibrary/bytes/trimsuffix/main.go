/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
	"os"
)

/*
rimSuffix返回s，不包含所提供的后缀字符串。如果s没有以后缀结尾，则返回不变的s。
*/
func main() {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!"), []byte("x!"))...)
	os.Stdout.Write(b)                                                      // Hello, world!
	fmt.Println(string(bytes.TrimSuffix([]byte("hi go go"), []byte("go")))) // hi go
}
