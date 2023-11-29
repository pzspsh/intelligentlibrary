/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:08:15
*/
package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
trimleftfuncc将s视为utf -8编码的字节，并通过切掉所有满足f(c)的前utf -8编码的代码点c来返回s的子切片。
*/
func main() {
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimLeftFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimLeftFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}

/*
Output:

-gopher
go-gopher!
go-gopher!567
*/
