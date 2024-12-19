/*
@File   : main.go
@Author : pan
@Time   : 2024-12-19 16:41:21
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 外层循环
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// 内层循环
			for j := 0; j < 100; j++ {
				wg.Add(1)
				go func(j int) {
					defer wg.Done()
					// 执行任务
					if j == 99 {
						fmt.Printf("Task %d-%d completed\n", i, j)
					}
				}(j)
			}
		}(i)
	}
	// 等待所有任务完成
	wg.Wait()
	fmt.Println("All tasks completed")
}
