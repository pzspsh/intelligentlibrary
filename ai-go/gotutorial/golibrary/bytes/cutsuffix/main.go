/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 14:28:55
*/
package main

import (
	"bytes"
	"fmt"
)

/*
CutSuffix返回s，但不包含所提供的结束后缀字节片，并报告是否找到该后缀。如果s没有以后缀结尾，
则CutSuffix返回s, false。如果suffix是空字节片，则CutSuffix返回s, true。

CutSuffix返回原始片s的片，而不是副本。
*/
func main() {
	show := func(s, sep string) {
		before, found := bytes.CutSuffix([]byte(s), []byte(sep))
		fmt.Printf("CutSuffix(%q, %q) = %q, %v\n", s, sep, before, found)
	}
	show("Gopher", "Go")
	show("Gopher", "er")
}
