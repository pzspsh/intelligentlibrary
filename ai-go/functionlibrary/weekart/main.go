/*
@File   : main.go
@Author : pan
@Time   : 2024-09-18 13:38:47
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

func WeekDay(date string) int { // 获取某年月日是星期几
	if date == "" {
		return 0
	}
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		loc = time.FixedZone("CST", 8*3600)
	}
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", date, loc)
	week := t.Weekday().String()
	if week == "Sunday" {
		return 7
	} else if week == "Saturday" {
		return 6
	} else if week == "Friday" {
		return 5
	} else if week == "Thursday" {
		return 4
	} else if week == "Wednesday" {
		return 3
	} else if week == "Tuesday" {
		return 2
	} else if week == "Monday" {
		return 1
	} else {
		return 0
	}
}

func GetWeekRange(dateStr string) (start, end time.Time, err error) { // 获取某日期所在周的日期范围
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	weekday := date.Weekday()
	offsetToMonday := int(time.Monday - weekday)
	if offsetToMonday > 0 {
		offsetToMonday -= 7
	}
	start = date.AddDate(0, 0, offsetToMonday)
	end = start.AddDate(0, 0, 6)
	return start, end, nil
}

func GetDaysInYear(value interface{}) int { // 获取某年的天数
	var year int
	switch v := value.(type) {
	case int:
		year = v
	case string:
		year, _ = strconv.Atoi(v)
	}
	if year == 1582 {
		return 355
	}
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) { // 判断是否是闰年
		return 366
	}
	return 365
}

func GetDaysInMonth(value string) int { // 获取某年某月的天数
	t, err := time.Parse("2006-01", value)
	if err != nil {
		return 0
	}
	year := t.Year()
	month := t.Month()
	if year == 1582 && month == 10 {
		return 21
	}
	nextMonth := month + 1
	if nextMonth > 12 {
		nextMonth = 1
		year++
	}
	date := time.Date(year, nextMonth, 0, 0, 0, 0, 0, time.UTC)
	return date.Day()
}

func GetWeekNumber(year int) map[int]map[int]map[string]string {
	weeknum := map[int]map[int]map[string]string{}
	for month := 1; month <= 12; month++ {
		date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
		for day := 1; day <= date.Day(); day++ {
			d := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			weekday := d.Weekday()
			offsetToMonday := int(time.Monday - weekday)
			if offsetToMonday > 0 {
				offsetToMonday -= 7
			}
			start := d.AddDate(0, 0, offsetToMonday)
			end := start.AddDate(0, 0, 6)
			year, week := d.ISOWeek()
			monday := start.Format("2006-01-02")
			sunday := end.Format("2006-01-02")
			if _, ok := weeknum[year]; !ok {
				weeknum[year] = map[int]map[string]string{week: {monday: sunday}}
			} else if _, ok := weeknum[year][week]; !ok {
				weeknum[year][week] = map[string]string{monday: sunday}
			}
		}
	}
	for year, values := range weeknum {
		reslen := len(values)
		if reslen >= 2 {
			for i := 0; i <= reslen; i++ {
				value := values[i]
				for key, value := range value {
					fmt.Printf("year: %v, week: %v, start: %v, end: %v\n", year, i, key, value)
				}
			}
		} else {
			for week, value := range values {
				for key, value := range value {
					fmt.Printf("year: %v, week: %v, start: %v, end: %v\n", year, week, key, value)
				}
			}
		}
	}
	return weeknum
}

func GetYearWeekNum(year int) int { // 获取每年的周数
	date := time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)
	yearday := date.YearDay()
	firstDayOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	weekdayoffirstday := firstDayOfYear.Weekday()
	isoweek := 1
	isoweekoffset := 0
	if weekdayoffirstday == time.Friday {
		isoweekoffset = 3
	} else if weekdayoffirstday == time.Saturday {
		isoweekoffset = 2
	} else if weekdayoffirstday == time.Sunday {
		isoweekoffset = 1
	} else {
		isoweekoffset = 4 - int(weekdayoffirstday)
	}
	isoweek += (yearday - 1 - isoweekoffset) / 7
	if isoweek < 1 {
		isoweek = 1
	}
	return isoweek
}

func GetYearWeekInNum(value string) int { // 获取每年的周数
	if value == "" {
		return 0
	}
	t, err := time.Parse("2006", value)
	if err != nil {
		return 0
	}
	year := t.Year()
	if year > 1582 {
		return GetYearWeekInNum(value)
	} else if year == 1582 {
		return 355 / 7
	} else if year < 1582 {
		return GetOldYearWeek(year)
	}
	return 0
}

func GetTodayWeek(datestr string) int { // 获取今天是今天的第几周
	date, _ := time.Parse("2006-01-02", datestr)
	_, week := date.ISOWeek()
	return week
}

func GetYearDay(datestr string) int { // 获取某天是某年的第几天
	date, _ := time.Parse("2006-01-02", datestr)
	return date.YearDay()
}

func GetOldYearWeek(year int) int {
	return 0
}

// 获取给定年份的第n周的日期范围
func GetWeekDateRange(year int, weekNum int) (start, end time.Time) {
	// 设置地点，因为不同地点的周起始日可能不同
	loc, _ := time.LoadLocation("Asia/Shanghai")
	// 构建给定年份的第一天
	firstDayOfYear := time.Date(year, 1, 1, 0, 0, 0, 0, loc)
	// 计算第一天是星期几（0是星期日，6是星期六）
	weekdayOfFirstDay := int(firstDayOfYear.Weekday())
	// 计算第一周的第一天（星期一）
	firstMonday := firstDayOfYear.AddDate(0, 0, -weekdayOfFirstDay+1)
	// 计算第n周的第一天
	startOfWeek := firstMonday.AddDate(0, 0, (weekNum-1)*7)
	// 计算第n周的最后一天
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	fmt.Printf("第%d周的日期范围是：%s 至 %s\n", weekNum, start.Format("2006-01-02"), end.Format("2006-01-02"))
	return startOfWeek, endOfWeek
}

func main() {
	// fmt.Println(355 / 7)
	// fmt.Println(GetYearWeekInNum("1582-10"))
	// fmt.Println(WeekDay("1580-01-01 11:57:02"))
	// fmt.Println(GetDaysInMonth("2024-09"))
	// fmt.Println(GetYearWeekNum(2023))
	/*
		dateStr := "1583-01-01"
		start, end, err := GetWeekRange(dateStr)
		if err != nil {
			fmt.Println("日期解析出错:", err)
			return
		}
		fmt.Printf("日期 %s 所在周的范围是 %s 到 %s\n", dateStr, start.Format("2006-01-02"), end.Format("2006-01-02"))
	*/
	/*
		yearstart := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
		fmt.Println(yearstart.ISOWeek())
	*/
	GetWeekNumber(2023)
	// 计算1582年10月15日之前是错误的，因为10月没有10天，实际应该为21
	// GetDaysInMonth(1583, 9)
	// 周一(1582.10.1)-周日(1582.10.17)
}
