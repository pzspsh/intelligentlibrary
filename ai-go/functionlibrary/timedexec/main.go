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

type TimeOption struct {
	Mode   string // daily、weekly、monthly、yearly
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Second string
}

// 定时器、定时执行
func ScheduleTaskDemo() {
	hour := 10   // 设置要执行任务的时间点，这里是14点
	minute := 46 // 设置要执行任务的时间，这里是每天的14点57分
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

func ScheduleTask(interval string, taskTime string, task func()) {

}

func SecondsTimer() { // 秒钟计时器
	for {
		time.Sleep(5 * time.Second)
		fmt.Println("五秒钟执行一次")
	}

}

func MinutesTimer() { // 分钟计时器
	for {
		time.Sleep(5 * time.Minute)
		fmt.Println("五分钟执行一次")
	}
}

func HourTimer() { // 小时计时器
	for {
		time.Sleep(5 * time.Hour)
		fmt.Println("五个小时执行一次")
	}
}

func DailyTimer() { // 每天计时器
	hour := 15
	minute := 44
	second := 0
	for {
		targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, time.Local)
		if time.Now().After(targetTime) {
			targetTime = targetTime.AddDate(0, 0, 1)
		}
		fmt.Println("下次执行时间:", targetTime)
		time.Sleep(time.Until(targetTime))
		fmt.Printf("现在是%v点%v分，执行任务！\n", hour, minute)
	}
}

func WeeklyTimer() { // 每周计时器
	hour := 13
	minute := 48
	second := 0
	mode := "Friday"
	for {
		weekday := ParseWeeklyTime(mode)
		days := int(weekday - time.Now().Weekday())
		if days <= 0 {
			days += 7
		}
		targetTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), hour, minute, second, 0, time.Local)
		if time.Now().After(targetTime) {
			targetTime = targetTime.AddDate(0, 0, days)
		}
		fmt.Println("下次执行时间:", targetTime)
		time.Sleep(time.Until(targetTime))
		fmt.Printf("现在是%v点%v分，执行任务！\n", hour, minute)
	}
}

func MonthlyTimer() { // 每月计时器
	hour := 14
	minute := 30
	day := 19
	second := 0
	for {
		targetTime := time.Date(time.Now().Year(), time.Now().Month(), day, hour, minute, second, 0, time.Local)
		if time.Now().After(targetTime) {
			targetTime = targetTime.AddDate(0, 1, 0)
		}
		fmt.Println("下次执行时间:", targetTime)
		time.Sleep(time.Until(targetTime))
		fmt.Printf("现在是%v点%v分，执行任务！\n", hour, minute)
	}
}

func YearlyTimer() { // 每年计时器
	month := 9
	day := 20
	hour := 15
	minute := 42
	second := 0
	for {
		targetTime := time.Date(time.Now().Year(), time.Month(month), day, hour, minute, second, 0, time.Local)
		if time.Now().After(targetTime) {
			targetTime = targetTime.AddDate(1, 0, 0)
		}
		fmt.Println("下次执行时间:", targetTime)
		time.Sleep(time.Until(targetTime))
		fmt.Printf("现在是%v点%v分，执行任务！\n", hour, minute)
	}
}

func ParseDailyTime() {

}

func ParseWeeklyTime(dayName string) time.Weekday {
	switch dayName {
	case "Sunday":
		return time.Sunday
	case "Monday":
		return time.Monday
	case "Tuesday":
		return time.Tuesday
	case "Wednesday":
		return time.Wednesday
	case "Thursday":
		return time.Thursday
	case "Friday":
		return time.Friday
	case "Saturday":
		return time.Saturday
	default:
		return time.Sunday
	}
}

func ParseMonthlyTime() {

}

func ParseYearlyTime() {

}

func main() {
	// var month int
	// var day int
	// t := "01-03 10:11:00"
	// fmt.Sscanf(t, "%02d-%02d %s", &month, &day, &t)
	// fmt.Println(month, day)
	// var day int
	// t := "01-02 15:04"
	// fmt.Sscanf(t, "%d %s", &day, &t)
	// fmt.Println(day)
	// var dayName string
	// t := "Monday"
	// fmt.Sscanf(t, "%s %s", &dayName, &t)
	// weekDay := parseWeekDay(dayName)
	// fmt.Println("BBBBBBBBB", int(weekDay))
	// ScheduleTask()
	DailyTimer()
	// WeeklyTimer()
	// MonthlyTimer()
	// YearlyTimer()
}
