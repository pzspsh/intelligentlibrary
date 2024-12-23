/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:08:40
*/
package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	s := "abc"
	clone := strings.Clone(s)
	fmt.Println(s == clone)
	fmt.Println(unsafe.StringData(s) == unsafe.StringData(clone))
}
