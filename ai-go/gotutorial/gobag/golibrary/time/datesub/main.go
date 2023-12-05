/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:13:36
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//获取西雅图时间
	loc, _ := time.LoadLocation("America/Los_Angeles")
	usTime := time.Now().In(loc)
	//计算与当地时间到今天0点的时间差
	todayZero := time.Date(usTime.Year(), usTime.Month(), usTime.Day(), 0, 0, 0, 0, loc)
	duration := todayZero.Sub(usTime)
	fmt.Printf("time diff: %s\n", duration)
}
