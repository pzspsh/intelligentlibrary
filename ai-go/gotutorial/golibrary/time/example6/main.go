/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 09:53:55
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 返回当前时间，注意此时返回的是 time.Time 类型
	now := time.Now()
	fmt.Println(now)
	// 当前时间戳
	fmt.Println(now.Unix())
	// 纳秒级时间戳
	fmt.Println(now.UnixNano())
	// 时间戳小数部分 单位：纳秒
	fmt.Println(now.Nanosecond())

	// 返回日期
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	// 年
	fmt.Println(now.Year())
	// 月
	fmt.Println(now.Month())
	// 日
	fmt.Println(now.Day())

	// 时分秒
	hour, minute, second := now.Clock()
	fmt.Printf("hour:%d, minute:%d, second:%d\n", hour, minute, second)
	// 时
	fmt.Println(now.Hour())
	// 分
	fmt.Println(now.Minute())
	// 秒
	fmt.Println(now.Second())

	// 返回星期
	fmt.Println(now.Weekday())
	//返回一年中对应的第几天
	fmt.Println(now.YearDay())
	//返回时区
	fmt.Println(now.Location())

	// 返回一年中第几天
	fmt.Println(now.YearDay())
	/*
	   Go 语言提供了时间类型格式化函数 Format()，需要注意的是 Go 语言格式化时间模板不是常见的 Y-m-d H:i:s，而是 2006-01-02 15:04:05，也很好记忆(2006 1 2 3 4 5)。
	*/
	fmt.Println(now.Format("2006-01-02 15:03:04"))
	fmt.Println(now.Format("2006-01-02"))
	fmt.Println(now.Format("15:03:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))

	layout := "2006-01-02 15:04:05"
	t := time.Unix(now.Unix(), 0) // 参数分别是：秒数,纳秒数
	fmt.Println(t.Format(layout))

	//根据指定时间返回 time.Time 类型
	//分别指定年，月，日，时，分，秒，纳秒，时区
	t1 := time.Date(2011, time.Month(3), 12, 15, 30, 20, 0, now.Location())
	fmt.Println(t1.Format(layout))

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(t2)
	// 输出 2021-01-10 17:28:50 +0800 CST
	// time.Local 指定本地时间

	fmt.Println(time.Now())
	fmt.Println(time.Now().Location())
	t3, _ := time.Parse("2006-01-02 15:04:05", "2021-01-10 15:01:02")
	fmt.Println(t3)

	fmt.Println(now)

	// 1小时1分1s之后
	t4, _ := time.ParseDuration("1h1m1s")
	fmt.Println(t1)
	m1 := now.Add(t4)
	fmt.Println(m1)

	// 1小时1分1s之前
	t5, _ := time.ParseDuration("-1h1m1s")
	m2 := now.Add(t5)
	fmt.Println(m2)

	// 3小时之前
	t6, _ := time.ParseDuration("-1h")
	m3 := now.Add(t6 * 3)
	fmt.Println(m3)

	// 10 分钟之后
	t7, _ := time.ParseDuration("10m")
	m4 := now.Add(t7)
	fmt.Println(m4)

	// Sub 计算两个时间差
	sub1 := now.Sub(m3)
	fmt.Println(sub1.Hours())   // 相差小时数
	fmt.Println(sub1.Minutes()) // 相差分钟数

	// 1小时之后
	t8, _ := time.ParseDuration("1h")
	m5 := now.Add(t8)
	fmt.Println(m5)

	fmt.Println(m5.After(now))
	fmt.Println(now.Before(m5))
	fmt.Println(now.Equal(m5))
}
