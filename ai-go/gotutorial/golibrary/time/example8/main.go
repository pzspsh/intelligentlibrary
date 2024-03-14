/*
@File   : main.go
@Author : pan
@Time   : 2024-03-14 16:21:35
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	specTime := "2024-03-14 17:00:00"
	layout := "2006-01-02 15:04:05"
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}
	begin, err := time.ParseInLocation(layout, specTime, loc)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	now := time.Now().In(loc)
	duration := now.Sub(begin)
	seconds := int(duration.Seconds())
	fmt.Println(duration, " aaaaaaaaa ", seconds)

}
