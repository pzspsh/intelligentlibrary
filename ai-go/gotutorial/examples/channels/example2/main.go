/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:49:04
*/
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)
	// 主线程阻塞在接收端，直至goroutine发送结束消息
	<-done
}
