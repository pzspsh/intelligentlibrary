/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 01:06:06
*/
package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	age      int
	language string
}

func main() {
	p := Programmer{"stefno", 18, "go"}
	fmt.Println(p)

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "Golang"

	fmt.Println(p)
}
