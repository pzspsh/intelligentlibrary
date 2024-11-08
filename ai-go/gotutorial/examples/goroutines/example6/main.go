/*
@File   : main.go
@Author : pan
@Time   : 2024-11-08 14:37:49
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineDemo() {
	var gonumber = make(chan int, 10)
	var count int = 100
	for i := 0; i < count; i++ {
		gonumber <- 1
		go func(c int, cn chan int) {
			fmt.Println("目标：", c)
			time.Sleep(1 * time.Second)
			<-cn
		}(i, gonumber)
	}
}

func GoroutineDemo1() {
	var wg sync.WaitGroup
	var count int = 100
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(c int) {
			defer wg.Done()
			fmt.Println("目标：", c)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}

func GoroutineDemo2() {
	var wg sync.WaitGroup
	var gonumber = make(chan int, 10)
	var count int = 100
	for i := 0; i < count; i++ {
		wg.Add(1)
		gonumber <- 1
		go func(c int, cn chan int) {
			defer wg.Done()
			fmt.Println("目标：", c)
			time.Sleep(1 * time.Second)
			<-cn
		}(i, gonumber)
	}
	wg.Wait()
}

func main() {
	GoroutineDemo1()
}
