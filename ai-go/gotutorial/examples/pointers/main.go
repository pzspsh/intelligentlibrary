/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:39:34
*/
package main

import (
	"fmt"
)

// 通过值传递，仅拷贝变量，不会修改实际值
func zeroval(ival int) {
	ival = 0
}

// 通过指针传递，可以修改实际值
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// 用&获取指针，即变量i的地址
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
