package main

import (
	"bytes"
	"fmt"
	"unicode"
)

/*
TrimFunc通过切掉所有满足f(c)的开头和结尾的utf -8编码的代码点c，返回s的子切片。
*/
func main() {
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("\"go-gopher!\""), unicode.IsLetter)))
	fmt.Println(string(bytes.TrimFunc([]byte("go-gopher!"), unicode.IsPunct)))
	fmt.Println(string(bytes.TrimFunc([]byte("1234go-gopher!567"), unicode.IsNumber)))
}

/*
Output:

-gopher!
"go-gopher!"
go-gopher
go-gopher!
*/
