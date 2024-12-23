/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:37:30
*/
package main

import (
	"context"
	"fmt"
	"time"
)

/*
这个例子传递了一个带有超时的上下文，告诉阻塞函数在超时过后应该放弃它的工作。
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
