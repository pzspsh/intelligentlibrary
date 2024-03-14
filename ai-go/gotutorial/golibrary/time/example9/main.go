/*
@File   : main.go
@Author : pan
@Time   : 2024-03-14 17:07:24
*/
package main

import (
	"fmt"
	"time"
)

/*
golang编写某个时间到现在时间是否大于一天
*/
func main() {
	layout := "2006-01-02 15:04:05"
	DemoDay1(layout, "2024-03-13 18:00:00")

}

func DemoDay1(layout, sptime string) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}
	begin, err := time.ParseInLocation(layout, sptime, loc)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	now := time.Now().In(loc)
	duration := now.Sub(begin)
	oneday := 24 * time.Hour
	if duration > oneday {
		fmt.Println("已经超过一天了")
	} else {
		fmt.Println("还没有超过一天")
	}
}

func DemoDay() {
	// 以 Unix 时间戳的形式定义一个时间点
	// 例如，这里是2023年8月16日12:00:00
	startTime := time.Unix(1660500800, 0)
	// 获取当前时间
	now := time.Now()
	// 计算两个时间点之间的持续时间
	duration := now.Sub(startTime)

	// 判断持续时间是否超过一天（24小时）
	if duration > 24*time.Hour {
		fmt.Println("已经超过一天了")
	} else {
		fmt.Println("还没有超过一天")
	}
}
