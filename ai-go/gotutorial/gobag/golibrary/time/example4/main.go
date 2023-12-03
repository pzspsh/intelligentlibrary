/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 02:11:09
*/
package main

import (
	"fmt"
	"time"
)

func run(task_id, sleeptime int, ch chan string) {
	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
}
func main() {
	input := []int{3, 2, 1}
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go run(i, sleeptime, chs[i])
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}
