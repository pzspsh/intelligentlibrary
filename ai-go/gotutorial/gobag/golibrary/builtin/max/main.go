/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:15:30
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// 比较整数
	fmt.Println(max(1, 2, 3)) // 3

	// 比较浮点数
	fmt.Println(max(1.5, 2.5, 3.5)) // 3.5

	// 浮点数包含 NaN
	fmt.Println(max(1.5, math.NaN(), 3.5)) // NaN

	// 比较字符串
	fmt.Println(max("apple", "banana", "cherry")) // cherry
}
