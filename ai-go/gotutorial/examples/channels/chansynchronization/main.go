/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:19:00
*/
package main

import (
	"fmt"
	"time"
)

// 这是一个我们将要在 Go 协程中运行的函数。`done` 通道
// 将被用于通知其他 Go 协程这个函数已经工作完毕。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// 发送一个值来通知我们已经完工啦。
	done <- true
}

func main() {
	// 运行一个 worker Go协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)
	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
}
