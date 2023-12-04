/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:16:12
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var sharedValue int64

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
				oldValue := atomic.LoadInt64(&sharedValue)
				newValue := oldValue + 1
				swapped := atomic.CompareAndSwapInt64(&sharedValue, oldValue, newValue)
				if swapped {
					fmt.Printf("Swapped: %d -> %d\n", oldValue, newValue)
				} else {
					fmt.Printf("Failed to swap: %d != %d\n", sharedValue, oldValue)
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Final shared value: %d\n", sharedValue)
}
