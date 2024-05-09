/*
@File   : main.go
@Author : pan
@Time   : 2024-05-09 13:38:23
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	tasknum := make(chan string, 5)
	go Receive(tasknum)
	Stat(tasknum)
}

func Receive(tasknum chan string) {
	for i := 0; i < 10; i++ {
		number := RandomNumber("RW", 6)
		tasknum <- number
		// time.Sleep(3 * time.Second)
	}
}

type TaskStat struct {
	End     bool
	Timeout time.Time
}

func Stat(tasknum chan string) {
	task := map[string]TaskStat{}
	for {
		if len(tasknum) > 0 {
			if rwnumber, ok := <-tasknum; ok {
				if _, ok := task[rwnumber]; !ok {
					task[rwnumber] = TaskStat{End: false, Timeout: time.Now()}
				}
			}
		}
		if len(task) > 0 {
			fmt.Println("循环更新task里面所有数据的状态")
			for rw, value := range task {
				fmt.Println("更新完这条rw的任务数据状态")
				fmt.Println("如果这个rw的任务数据状态时间太长，则从task中删除rw")
				now := time.Now()
				if now.After(value.Timeout.Add(5 * time.Second)) {
					fmt.Println("超时", rw, now, value.Timeout)
					delete(task, rw) // 删除时task会改变，所以需要重新获取task
				}
			}
		}
		fmt.Println("等待10秒后继续")
		time.Sleep(1 * time.Second)
	}
}

const (
	charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomNumber(prefix string, count int) string {
	sb := strings.Builder{}
	sb.Grow(count)
	for i := 0; i < count; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	number := fmt.Sprintf("%v-%v-%v", prefix, time.Now().Unix(), sb.String())
	return number
}
