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

func StatsRun() {
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

func StatsRun2() {
	var flag = make(chan string, 1)
	var updatechan = make(chan bool, 2)
	var snapshot = make(map[string]bool, 20)
	var taskmap = make(map[string]TaskStat, 20)
	go func() {
		for {
			tasks := []string{RandomStr(10), RandomStr(10), "task3", "task1", "task2", "task3", "task4"}
			for _, task := range tasks {
				mu.Lock()
				if _, ok := taskmap[task]; !ok {
					taskmap[task] = TaskStat{End: false}
					flag <- task
				}
				mu.Unlock()
			}
			time.Sleep(100 * time.Microsecond)
		}
	}()

	go func() {
		for {
			if rwnumber, ok := <-flag; ok {
				mu.Lock()
				snapshot[rwnumber] = true
				mu.Unlock()
			}
		}
	}()
	for {
		if len(snapshot) > 0 {
			mu.Lock()
			for rwnumber := range snapshot {
				updatechan <- true
				go func(update chan bool, key string) {
					defer func() { <-update }()
					mu.Lock()
					if value, exist := taskmap[key]; exist {
						fmt.Println(key, value)
						delete(taskmap, key)
						delete(snapshot, key)
					}
					time.Sleep(300 * time.Microsecond)
					mu.Unlock()
				}(updatechan, rwnumber)
			}
			mu.Unlock()
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	StatsRun2()
}
