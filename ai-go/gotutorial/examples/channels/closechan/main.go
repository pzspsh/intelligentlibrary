/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:52:22
*/
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 关闭channel且读完通道中所有消息后，more取到false，结束通信
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 用close函数关闭channel
	close(jobs)
	fmt.Println("sent all jobs")

	// 用消息阻塞保证goroutine先于主线程执行完
	<-done
}
