/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 17:12:30
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 2
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*4)
	defer cancelFunc()
	// 3
	go targetFunc(ctx)

	/*
		go func() {
			time.Sleep(time.Second * 1)
			cancelFunc()
		}()
	*/
	// 4
	for {
		select {
		case <-ctx.Done():
			// 5
			switch ctx.Err() {
			case context.DeadlineExceeded:
				fmt.Println("context timeout exceeded")
				return
			case context.Canceled:
				fmt.Println("context cancelled by force")
				return
			}
		default:
			time.Sleep(time.Second * 1)
			fmt.Println("sleep 1s")
		}
	}
}

// 1
func targetFunc(ctx context.Context) {
	time.Sleep(time.Second * 3)
	fmt.Println("u r here.")
}
