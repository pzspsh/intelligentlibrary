/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 11:54:27
*/
package main

import (
	"context"
	"fmt"
	"time"
)

/*
让我们从写一个简单的例子开始，我们有一个函数，它有一个从main调用的长时间运行代码。当调用方发出取消信号时，
必须终止长时间运行的函数。让我们讨论一下如何在上下文的帮助下实现这一点。
*/
func longRunning(ctx context.Context) (int, error) {
	count := 0
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			count = count + i
			fmt.Println("Current value of count:", count)
			time.Sleep(2 * time.Second)
		}
	}
	return count, nil
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		cancelFunc()
	}()
	count, err := longRunning(ctx)
	if err != nil {
		fmt.Println("long running task exited with error", err)
		return
	}

	fmt.Println("count is ", count)
}
