/*
@File   : main.go
@Author : pan
@Time   : 2024-07-22 15:07:55
*/
package main

import (
	"fmt"
)

type MyStruct struct {
	Field1 string
	Field2 int
	Field3 bool
}

func main() {
	var s1 *MyStruct  // nil指针
	s2 := &MyStruct{} // 指向包含零值字段的结构体的指针

	fmt.Println(s1 == nil) // 输出: true
	fmt.Println(s2 == nil) // 输出: false
}
