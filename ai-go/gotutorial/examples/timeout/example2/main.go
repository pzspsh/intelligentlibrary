package main

import (
	"fmt"
	"time"
)

func main() {
	// 将字符串形式的时间转换为Go中的时间类型
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// begin := time.Now()
	// fmt.Println(begin)
	begin, err := time.ParseInLocation(layout, time.Now().Format(layout), loc)
	if err != nil {
		panic(err)
	}
	// 获取当前时间
	time.Sleep(5 * time.Second)
	now := time.Now().In(loc)

	// 计算时间差
	duration := now.Sub(begin)

	// 将duration转换为时分秒
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	fmt.Printf("相差 %02d:%02d:%02d\n", hours, minutes, seconds)
}
