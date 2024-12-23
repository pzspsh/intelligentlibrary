/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:14:47
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// 比较整数
	fmt.Println(min(1, 2, 3)) // 1
	// 比较浮点数
	fmt.Println(min(1.5, 2.5, 3.5)) // 1.5
	// 浮点数包含 NaN
	fmt.Println(min(1.5, math.NaN(), 3.5)) // NaN
	// 比较字符串
	fmt.Println(min("apple", "banana", "cherry")) // apple
}
