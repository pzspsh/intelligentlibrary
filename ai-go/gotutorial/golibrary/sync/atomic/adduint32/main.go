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

var counter uint32

func main() {
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint32(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter value:", counter)
}
