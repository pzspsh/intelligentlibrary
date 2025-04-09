/*
@File   : main.go
@Author : pan
@Time   : 2023-12-03 00:05:14
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", difference)

	startTime := time.Now() // 开始时间
	// 假设我们做一些事情，然后获取结束时间
	time.Sleep(3 * time.Second)                  // 模拟等待
	endTime := time.Now()                        // 结束时间
	duration := endTime.Sub(startTime).Seconds() // 计算时间差
	fmt.Println("Time difference:", int64(duration))
}
