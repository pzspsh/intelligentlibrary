/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:56:28
*/
package main

import (
	"fmt"
	"time"
)

// worker从jobs通道获取任务，并将执行结果发送到results通道
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

/*
工人池（协程池）用于使用固定数量的协程处理任务，防止协程数量太多。在部分特殊场景需要使用。协程池并非Go语言原生概念，
而是基于已有语言特性搭建的开发模型。
*/
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// 开启worker数量为3的协程池
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 将新任务添加到jobs通道中
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
