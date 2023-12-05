/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:53:40
*/
package main

import (
	"fmt"
	"time"
)

/* 定时器用于在将来的某个时间点执行代码。通过从通道接收到消息的方式表示达到触发时间点，接收到的消息为当前时间。 */
func main() {
	// 用time.NewTimer()定义2秒后的定时器
	timer1 := time.NewTimer(2 * time.Second)
	// 阻塞在这里，直到到达触发时间点
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	// 用goroutine的方式避免阻塞主线程
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// 用Stop()停止定时器
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
