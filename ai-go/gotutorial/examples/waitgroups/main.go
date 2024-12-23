/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:57:25
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		// 对每个goroutine，等待计数加1
		wg.Add(1)
		i := i
		go func() {
			// 延迟执行，完成工作后，将等待计数减1
			defer wg.Done()
			worker(i)
		}()
	}
	// 当等待计数为0时，结束等待
	wg.Wait()
}
