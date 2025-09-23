package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RandomStr(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz1234567890")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenNumber() []string {
	var number []string
	for range 20 {
		number = append(number, RandomStr(10))
	}
	return number
}

func StatsRun() {
	var updatech = make(chan struct{}, 5)
	var taskchan = make(chan string, 10)
	var taskmap = &sync.Map{}
	go func() {
		for {
			numberlist := []string{RandomStr(10), RandomStr(10), "task3", "task1", "task2", "task3", "task4"}
			if len(numberlist) > 0 {
				fmt.Println("开始生成任务", numberlist)
				for _, number := range numberlist {
					if _, ok := taskmap.Load(number); !ok {
						taskchan <- number
						taskmap.Store(number, true)
					}
				}
				time.Sleep(time.Second * 1)
			} else {
				time.Sleep(time.Second * 10)
				continue
			}
		}
	}()
	for {
		if number, ok := <-taskchan; ok {
			updatech <- struct{}{}
			go func(number string, update chan struct{}) {
				defer func() { <-update }()
				time.Sleep(time.Second * 5)
				fmt.Println(number)
				taskmap.Delete(number)
			}(number, updatech)
		}
	}
}

func main() {
	StatsRun()
}
