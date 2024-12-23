/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:10:53
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
TrimRightFunc通过切掉所有满足f(c)的尾随utf -8编码的代码点c返回s的子切片。
*/
func main() {
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimRightFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}

/*
Output:

go-
go-gopher
1234go-gopher!
*/
