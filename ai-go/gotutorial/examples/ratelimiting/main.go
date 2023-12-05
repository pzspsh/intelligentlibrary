/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:58:33
*/
package main

import (
	"fmt"
	"time"
)

/* 速率限制是互联网和软件工程中的概念，用于防止对资源的使用超过一定限度。Go通过周期定时器和通道可以方便地实现这一功能。 */
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		// 每200毫秒结束阻塞，接收一次请求
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)
	// 先允许执行三次
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// 然后每200毫秒允许执行一次
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
