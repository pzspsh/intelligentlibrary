/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:37:30
*/
package main

import (
	"context"
	"fmt"
)

/*
此示例演示了如何使用可取消上下文来防止运行例程泄漏。在示例函数结束时，由gen启动的go例程将无泄漏地返回。
*/
func main() {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
