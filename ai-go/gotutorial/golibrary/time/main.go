/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:24:08
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	// 获取当前时间
	now := time.Now()
	p(now)

	// 创建一个指定的时间
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// 打印出时间的某项属性
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 注意：两个时间计算差值，返回的是duration
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
