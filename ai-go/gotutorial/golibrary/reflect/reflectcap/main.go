/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:30:02
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3}

	// 初始容量为3
	fmt.Println("Slice len: ", len(s), ", cap: ", cap(s))

	// 使用reflect修改容量
	v := reflect.ValueOf(&s).Elem()
	v.SetCap(5)

	// 容量变为5
	fmt.Println("Slice len: ", len(s), ", cap: ", cap(s))

	// 再次使用reflect修改容量
	v.SetCap(2)

	// 容量变为2，切片会被缩短
	fmt.Println("Slice len: ", len(s), ", cap: ", cap(s))
}
