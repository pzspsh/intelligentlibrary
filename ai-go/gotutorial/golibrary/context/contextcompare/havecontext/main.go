/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:08:15
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有截止时间的 context
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()
	// 启动一个 goroutine 来处理任务
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 如果 context 被取消了，就退出循环
				fmt.Println("goroutine 1: context canceled")
				return
			default:
				// 模拟一些耗时的操作，普通情况可能是rpc调用
				time.Sleep(1 * time.Second)
				fmt.Println("goroutine 1: doing some work")
			}
		}
	}(ctx)
	// 启动另一个 goroutine 来处理任务
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 如果 context 被取消了，就退出循环
				fmt.Println("goroutine 2: context canceled")
				return
			default:
				// 模拟一些耗时的操作
				time.Sleep(1 * time.Second)
				fmt.Println("goroutine 2: doing some work")
			}
		}
	}(ctx)
	// 等待一段时间，然后取消 context
	time.Sleep(3 * time.Second)
	cancel()
	fmt.Println("main: context canceled")
	time.Sleep(1 * time.Second)
}
