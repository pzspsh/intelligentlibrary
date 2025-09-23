package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var wg sync.WaitGroup

func main() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go increment()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}

func increment() {
	defer wg.Done()
	atomic.AddInt32(&counter, 1)
}
