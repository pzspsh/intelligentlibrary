/*
@File   : leakydrum.go
@Author : pan
@Time   : 2023-09-20 15:13:33
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

// 每个请求来了，把需要执行的业务逻辑封装成Task，放入木桶，等待worker取出执行
type Task struct {
	handler func() Result // worker从木桶中取出请求对象后要执行的业务逻辑函数
	resChan chan Result   // 等待worker执行并返回结果的channel
	taskID  int
}

// 封装业务逻辑的执行结果
type Result struct {
}

// 模拟业务逻辑的函数
func handler() Result {
	time.Sleep(300 * time.Millisecond)
	return Result{}
}

func NewTask(id int) Task {
	return Task{
		handler: handler,
		resChan: make(chan Result),
		taskID:  id,
	}
}

// 漏桶
type LeakyBucket struct {
	BucketSize int       // 木桶的大小
	NumWorker  int       // 同时从木桶中获取任务执行的worker数量
	bucket     chan Task // 存方任务的木桶
}

func NewLeakyBucket(bucketSize int, numWorker int) *LeakyBucket {
	return &LeakyBucket{
		BucketSize: bucketSize,
		NumWorker:  numWorker,
		bucket:     make(chan Task, bucketSize),
	}
}

func (b *LeakyBucket) validate(task Task) bool {
	// 如果木桶已经满了，返回false
	select {
	case b.bucket <- task:
	default:
		fmt.Printf("request[id=%d] is refused\n", task.taskID)
		return false
	}

	// 等待worker执行
	<-task.resChan
	fmt.Printf("request[id=%d] is run\n", task.taskID)
	return true
}

func (b *LeakyBucket) Start() {
	// 开启worker从木桶拉取任务执行
	go func() {
		for i := 0; i < b.NumWorker; i++ {
			go func() {
				for {
					task := <-b.bucket
					result := task.handler()
					task.resChan <- result
				}
			}()
		}
	}()
}

func main() {
	bucket := NewLeakyBucket(10, 4)
	bucket.Start()

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task := NewTask(id)
			bucket.validate(task)
		}(i)
	}
	wg.Wait()
}
