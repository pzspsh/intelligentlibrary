/*
@File   : main.go
@Author : pan
@Time   : 2024-08-02 16:39:40
*/
package main

import (
	"fmt"
	"time"
)

const (
	timezone = "Asia/Shanghai"
	layout   = "2006-01-02 15:04:05"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	start := time.Now().In(loc)
	time.Sleep(60 * time.Second)
	end := time.Now().In(loc)
	res := UsedTime(start, end)
	fmt.Println(res)
}

func UsedTime(start, end time.Time) string {
	loc, _ := time.LoadLocation(timezone)
	start, _ = time.ParseInLocation(layout, start.Format(layout), loc)
	duration := end.Sub(start)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
