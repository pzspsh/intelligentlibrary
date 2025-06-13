/*
@File   : main.go
@Author : pan
@Time   : 2025-06-03 15:42:04
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	mu sync.Mutex
)

type TaskStat struct {
	End     bool
	Timeout time.Time
}

func RandomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	var updatechan = make(chan bool, 2)
	taskmap := make(map[string]TaskStat, 20)
	go func() {
		for {
			tasks := []string{RandomStr(10), RandomStr(10), "task3", "task1", "task2", "task3", "task4"}
			for _, task := range tasks {
				mu.Lock()
				if _, ok := taskmap[task]; !ok {
					taskmap[task] = TaskStat{End: false}
				}
				mu.Unlock()
			}
			time.Sleep(100 * time.Microsecond)
		}
	}()
	for {
		if len(taskmap) > 0 {
			mu.Lock()
			snapshot := make([]string, 0, len(taskmap))
			for rwnumber := range taskmap {
				snapshot = append(snapshot, rwnumber)
			}
			mu.Unlock()

			for _, rwnumber := range snapshot {
				updatechan <- true
				go func(update chan bool, key string) {
					defer func() { <-update }()
					mu.Lock()
					if value, exist := taskmap[key]; exist {
						fmt.Println(key, value)
						delete(taskmap, key)
					}
					time.Sleep(300 * time.Microsecond)
					mu.Unlock()
				}(updatechan, rwnumber)
			}
		}
		fmt.Println("======================================================")
		time.Sleep(100 * time.Microsecond)
	}
}
