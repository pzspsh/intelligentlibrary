/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:54:14
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义了一个周期为500毫秒的周期定时器
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			// 每间隔500毫秒从定时器通道接收到消息
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 1.6秒后停止定时器
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
