/*
@File   : main.go
@Author : pan
@Time   : 2024-07-31 15:08:20
*/
package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex
	queue []interface{}
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(data interface{}) {
	q.Lock()
	defer q.Unlock()
	q.queue = append(q.queue, data)
}

func (q *Queue) Dequeue() interface{} {
	q.Lock()
	defer q.Unlock()
	if len(q.queue) == 0 {
		return nil
	}
	firstElement := q.queue[0]
	q.queue = q.queue[1:]
	return firstElement
}

func (q *Queue) Size() int {
	q.Lock()
	defer q.Unlock()
	return len(q.queue)
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Size()) // 输出: 3

	firstElement := queue.Dequeue()
	fmt.Println(firstElement) // 输出: 1
	fmt.Println(queue.Size()) // 输出: 2
}

/*
在这个示例中，我们定义了一个Queue结构体，它有一个互斥锁Mutex，一个切片queue用来存储元素。Enqueue方法用于添加元素到队列，
Dequeue方法用于从队列中移除元素，Size方法返回队列的大小。

这个示例展示了如何使用Go语言的基本功能实现一个简单的线程安全队列。这个队列是先进先出（FIFO）的，
但是如果你需要先进后出（LIFO）的队列，你可以使用Go语言标准库中的容器列表list实现，或者使用堆栈（Stack）
*/
