/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:53:22
*/
package main

import (
	"bytes"
	"fmt"
)

// TrimSpace通过切掉Unicode定义的所有前后空白返回s的子切片。
func main() {
	fmt.Printf("%s", bytes.TrimSpace([]byte(" \t\n a lone gopher \n\t\r\n"))) // a lone gopher
	fmt.Println(string(bytes.TrimSpace([]byte(" hello my bro "))))            // hello my bro
}
