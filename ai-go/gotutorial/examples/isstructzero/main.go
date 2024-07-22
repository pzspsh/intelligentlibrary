/*
@File   : main.go
@Author : pan
@Time   : 2024-07-22 15:07:36
*/
package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 string
	Field2 int
	Field3 bool
}

func isStructZero(s *MyStruct) bool {
	return reflect.DeepEqual(s, &MyStruct{})
}

func main() {
	var s1 *MyStruct                     // nil指针
	s2 := &MyStruct{}                    // 指向包含零值字段的结构体的指针
	s3 := &MyStruct{Field1: "not empty"} // 指向包含非零值字段的结构体的指针

	// 注意：不能直接对nil指针使用isStructZero，因为它会导致panic
	// 先检查指针是否为nil
	if s1 == nil {
		fmt.Println("s1 is nil")
	} else {
		fmt.Println("s1 is not empty:", isStructZero(s1)) // 输出: s1 is not empty: true（如果s1不是nil）
	}
	fmt.Println("s2 is empty:", isStructZero(s2))     // 输出: s2 is empty: true
	fmt.Println("s3 is not empty:", isStructZero(s3)) // 输出: s3 is not empty: false
}
