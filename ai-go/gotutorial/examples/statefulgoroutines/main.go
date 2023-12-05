/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:06:45
*/
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

/*
比起互斥锁，Go更推荐使用通道通信的方式来实现goroutine之间的同步。以下是一个使用channel通信来实现并发安全的缓存服务
的例子。缓存服务从reads和writes两个channel中读取读写命令，并依次执行，从而避免了多个goroutine对缓存数据的竞争。两个channel
在这里起到了消息队列的作用。
*/
func main() {
	var readOps uint64
	var writeOps uint64
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			//  缓存服务从两个channel取命令，串行执行，每次一条
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				// 接收到返回后将计数加1
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	// 用LoadUint64保证读取计数为原子操作
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
