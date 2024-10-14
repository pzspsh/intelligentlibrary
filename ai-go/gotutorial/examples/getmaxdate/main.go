/*
@File   : main.go
@Author : pan
@Time   : 2024-10-11 17:43:05
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 假设我们有一个包含10个时间日期的数组
	dates := []time.Time{
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
		// ...
		time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
	}

	// 使用内置的sort函数和自定义比较器来找到最大的时间
	maxDate := getMaxDate(dates)
	fmt.Println("最大的时间是:", maxDate)
}

func getMaxDate(dates []time.Time) time.Time {
	if len(dates) == 0 {
		return time.Time{}
	}
	max := dates[0]
	for _, t := range dates {
		if t.After(max) {
			max = t
		}
	}
	return max
}
