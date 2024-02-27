/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:49:18
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个父级 Context
	parentCtx := context.Background()

	// 创建一个带有取消函数的子级 Context，设置超时时间为2秒
	childCtx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel()

	// 在子级 Context 中启动一个 goroutine
	go doSomething(childCtx)

	// 阻塞等待一段时间
	time.Sleep(3 * time.Second)

	// 调用取消函数，取消子级 Context 中的操作
	cancel()

	// 阻塞等待一段时间
	time.Sleep(1 * time.Second)
}

func doSomething(ctx context.Context) {
	// 模拟一个耗时操作
	for {
		select {
		case <-ctx.Done():
			// 当 Context 被取消时，停止操作
			fmt.Println("Operation cancelled.")
			return
		default:
			// 模拟耗时操作
			fmt.Println("Performing operation...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
