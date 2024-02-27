/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:11:12
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Duration(10) * time.Minute
	// 结果为 10，10 正好是 1 的整数倍
	fmt.Println(a.Truncate(time.Duration(1) * time.Minute))
	// 结果为 9，3 的 3 倍是 9，最接近 10
	fmt.Println(a.Truncate(time.Duration(3) * time.Minute))
	// 结果为 8，4 的 2 倍是 8，最接近 10
	fmt.Println(a.Truncate(time.Duration(4) * time.Minute))
	// 结果为 10，10 正好是 5 的整数倍
	fmt.Println(a.Truncate(time.Duration(5) * time.Minute))
	// 结果为 6，6 的 1 倍是 6，最接近 10
	fmt.Println(a.Truncate(time.Duration(6) * time.Minute))
	// 结果为 0，11 的 1 倍是11，大于了 10
	fmt.Println(a.Truncate(time.Duration(11) * time.Minute))
}
