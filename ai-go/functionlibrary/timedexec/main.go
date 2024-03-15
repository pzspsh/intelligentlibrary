/*
@File   : main.go
@Author : pan
@Time   : 2024-03-14 17:35:59
*/
package main

import (
	"fmt"
	"time"
)

func scheduleTask() {
	hour := 15   // 设置要执行任务的时间点，这里是14点
	minute := 10 // 设置要执行任务的时间，这里是每天的14点57分
	second := 0
	for { // 主循环，持续检查并执行任务
		targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, time.Local)
		if time.Now().After(targetTime) { // 如果当前时间已经超过了目标时间，则设置目标时间为明天的14点30分
			targetTime = targetTime.AddDate(0, 0, 1)
		}
		time.Sleep(time.Until(targetTime))            // 等待到目标时间
		fmt.Printf("现在是%v点%v分，执行任务！\n", hour, minute) // 执行任务
	}
}

func main() {
	scheduleTask()
}
