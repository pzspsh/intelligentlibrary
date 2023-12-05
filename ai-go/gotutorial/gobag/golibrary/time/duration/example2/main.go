/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:10:27
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.Duration(10) * time.Minute
	// 结果为 10，10 正好是 1 的整数倍
	fmt.Println(a.Round(time.Duration(1) * time.Minute))
	// 结果为 9，3 的 3 倍是 9，3 的 4 倍是 12，9 和 10 离得更近
	fmt.Println(a.Round(time.Duration(3) * time.Minute))
	// 结果为 12，4 的 2 倍是 8，4 的 3 倍是 12，距离一样近往上入
	fmt.Println(a.Round(time.Duration(4) * time.Minute))
	// 结果为 10，10 正好是 5 的整数倍
	fmt.Println(a.Round(time.Duration(5) * time.Minute))
	// 结果为12， 6 的 1 倍是 6，6 的 2 倍是 12，10 离 12 更近
	fmt.Println(a.Round(time.Duration(6) * time.Minute))
}
