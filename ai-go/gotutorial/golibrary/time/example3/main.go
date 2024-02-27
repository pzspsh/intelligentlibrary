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
	ch := make(chan string)
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		go run(i, sleeptime, ch)
	}
	for range input {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))
}
