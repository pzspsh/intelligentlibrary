/*
@File   : main.go
@Author : pan
@Time   : 2024-10-11 17:26:44
*/
package main

import (
	"fmt"
	"time"
)

func TimeCompare() {
	// 定义两个时间
	time1, _ := time.Parse("2006-01-02", "2023-03-10")
	time2, _ := time.Parse("2006-01-02", "2023-03-15")
	// 比较时间
	if time1.Before(time2) {
		fmt.Println("Time1在time2之前") // Time1在time2之前
	} else if time1.After(time2) {
		fmt.Println("Time1在time2之后")
	} else if time1.Equal(time2) {
		fmt.Println("Time1等于time2")
	}
}

func TimeCompare1() {
	now := time.Now()
	yesterday := time.Now().AddDate(0, 0, -1)
	nowstart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	yesterdaystart := time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, time.Local)
	nowformated := nowstart.Format("2006-01-02 15:04:05")
	yesterdayFormatted := yesterdaystart.Format("2006-01-02 15:04:05")
	if nowstart.After(yesterdaystart) {
		fmt.Println(nowstart)
	} else {
		fmt.Println(yesterdaystart)
	}
	fmt.Println("yesterday: ", yesterday.UTC())
	fmt.Println("now: ", now.UTC())

	fmt.Println(nowformated)
	fmt.Println(yesterdayFormatted)
}

func main() {
	TimeCompare()
}
