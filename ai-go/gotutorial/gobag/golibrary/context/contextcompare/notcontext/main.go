/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:07:15
*/
package main

import (
	"fmt"
	"time"
)

// 全局变量，用于存储上下文信息
var (
	deadline time.Time
	// requestID string
)

func main() {
	// 设置上下文信息
	deadline = time.Now().Add(5 * time.Second)
	// requestID = "123456"
	// 启动一个 goroutine 来处理任务
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				// 模拟一些耗时的操作
				fmt.Println("goroutine 1: doing some work")
			default:
				// 检查上下文信息，如果已经超时或被取消了，就退出循环
				if time.Now().After(deadline) {
					fmt.Println("goroutine 1: context canceled")
					return
				}
			}
		}
	}()

	// 启动另一个 goroutine 来处理任务
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				// 模拟一些耗时的操作
				fmt.Println("goroutine 2: doing some work")
			default:
				// 检查上下文信息，如果已经超时或被取消了，就退出循环
				if time.Now().After(deadline) {
					fmt.Println("goroutine 2: context canceled")
					return
				}
			}
		}
	}()

	// 等待一段时间，然后取消上下文信息
	time.Sleep(3 * time.Second)
	fmt.Println("main: context canceled")
	deadline = time.Now()
	time.Sleep(1 * time.Second)
}
