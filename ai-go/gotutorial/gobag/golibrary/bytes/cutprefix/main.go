/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 16:59:29
*/
package main

import (
	"bytes"
	"fmt"
)

/*
CutPrefix返回s，但未提供前导前缀字节切片并报告是否找到前缀。如果s不以前缀
开头，则CutPrefix返回s，false。如果prefix是空字节片，则CutPrefix返回s，true。

CutPrefix 返回原始切片的切片，而不是副本。
*/
func main() {
	show := func(s, sep string) {
		after, found := bytes.CutPrefix([]byte(s), []byte(sep))
		fmt.Printf("CutPrefix(%q, %q) = %q, %v\n", s, sep, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
}
