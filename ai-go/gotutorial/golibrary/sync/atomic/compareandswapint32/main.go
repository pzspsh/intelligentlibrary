/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:16:12
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				incrementCounter()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func incrementCounter() {
	for {
		current := atomic.LoadInt32(&counter)
		new := current + 1
		if atomic.CompareAndSwapInt32(&counter, current, new) {
			break
		}
	}
}
