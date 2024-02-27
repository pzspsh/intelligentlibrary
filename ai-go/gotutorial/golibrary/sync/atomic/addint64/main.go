/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:11:33
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup
	// Number of goroutines
	numGoroutines := 100
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("Counter value: %d\n", counter)
}
