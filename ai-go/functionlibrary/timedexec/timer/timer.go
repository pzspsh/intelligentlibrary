/*
@File   : timer.go
@Author : pan
@Time   : 2024-09-20 13:52:50
*/
package timer

import (
	"fmt"
	"time"
)

func ScheduleTask(interval string, taskTime string, task func()) {
	switch interval {
	case "daily":
		StartDailyTimer(taskTime, task)
	case "weekly":
		StartWeeklyTimer(taskTime, task)
	case "monthly":
		StartMonthlyTimer(taskTime, task)
	case "yearly":
		StartYearlyTimer(taskTime, task)
	default:
		fmt.Printf("interval value %v not exist", interval)
	}
}

func StartDailyTimer(targetTime string, task func()) {
	for {
		now := time.Now()
		target := getNextTime(targetTime, now, "daily")
		time.Sleep(time.Until(target))
		task()
	}
}

func StartWeeklyTimer(targetTime string, task func()) {
	for {
		now := time.Now()
		target := getNextTime(targetTime, now, "weekly")
		time.Sleep(time.Until(target))
		task()
	}
}

func StartMonthlyTimer(targetTime string, task func()) {
	for {
		now := time.Now()
		target := getNextTime(targetTime, now, "monthly")
		time.Sleep(time.Until(target))
		task()
	}
}

func StartYearlyTimer(targetTime string, task func()) {
	for {
		now := time.Now()
		target := getNextTime(targetTime, now, "yearly")
		time.Sleep(time.Until(target))
		task()
	}
}

func getNextTime(targettime string, now time.Time, mode string) time.Time {
	year, month, day := now.Date()
	switch mode {
	case "daily":
		hour, min, sec := parseTime(targettime)
		targetTime := time.Date(year, month, day, hour, min, sec, 0, time.Local)
		if now.After(targetTime) {
			targetTime = targetTime.AddDate(0, 0, 1)
		}
		return targetTime
	case "weekly":
		weekday, t := parseWeeklyTime(targettime)
		hour, min, sec := parseTime(t)
		days := int(weekday - now.Weekday())
		if days <= 0 {
			days += 7
		}
		targetTime := time.Date(year, month, day, hour, min, sec, 0, time.Local)
		if now.After(targetTime) {
			targetTime = targetTime.AddDate(0, 0, days)
		}
		return targetTime
	case "monthly":
		days, t := parseMonthlyTime(targettime)
		hour, min, sec := parseTime(t)
		targetTime := time.Date(year, month, days, hour, min, sec, 0, time.Local)
		if now.After(targetTime) {
			targetTime = targetTime.AddDate(0, 1, 0)
		}
		return targetTime
	case "yearly":
		month, day, t := parseYearlyTime(targettime)
		hour, min, sec := parseTime(t)
		targetTime := time.Date(year, month, day, hour, min, sec, 0, time.Local)
		if now.After(targetTime) {
			targetTime = targetTime.AddDate(1, 0, 0)
		}
		return targetTime
	default:
		return now
	}
}

func parseTime(t string) (int, int, int) {
	hour, min, sec := 0, 0, 0
	fmt.Sscanf(t, "%d:%d:%d", &hour, &min, &sec)
	return hour, min, sec
}

func parseWeeklyTime(t string) (time.Weekday, string) {
	var dayname string
	fmt.Sscanf(t, "%s %s", &dayname, &t)
	weekday := parseWeekDay(dayname)
	return weekday, t
}

func parseWeekDay(dayName string) time.Weekday {
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

func parseMonthlyTime(t string) (int, string) {
	var day int
	fmt.Sscanf(t, "%d %s", &day, &t)
	return day, t
}

func parseYearlyTime(t string) (time.Month, int, string) {
	var month int
	var day int
	fmt.Sscanf(t, "%02d-%02d %s", &month, &day, &t)
	return time.Month(month), day, t
}

func SecondsTimer(sec int, task func()) { // 秒钟计时器
	for {
		time.Sleep(time.Duration(sec) * time.Second)
		task()
	}

}

func MinutesTimer(min int, task func()) { // 分钟计时器
	for {
		time.Sleep(time.Duration(min) * time.Minute)
		task()
	}
}

func HourTimer(hour int, task func()) { // 小时计时器
	for {
		time.Sleep(time.Duration(hour) * time.Hour)
		task()
	}
}
