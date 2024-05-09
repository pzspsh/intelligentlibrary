/*
@File   : main.go
@Author : pan
@Time   : 2024-05-09 16:58:19
*/
package main

import (
	"fmt"
	"time"
)

type TaskStat struct {
	End     bool
	Timeout time.Time
}

func main() {
	task := map[string]TaskStat{}
	start := "开始"
	for i := 0; i < 10; i++ {
		starts := fmt.Sprintf("%s%d", start, i)
		task[starts] = TaskStat{End: false, Timeout: time.Now()}
	}
	for key, value := range task {
		fmt.Println("删除", key, value)
		delete(task, key)
	}
	fmt.Println("BBBB", task)
}
