/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:25:07
*/
package main

import (
	"fmt"

	"sync"

	"sync/atomic"

	"time"
)

func main() {

	var counter uint64 = 0

	var wg sync.WaitGroup

	// Writer goroutine
	wg.Add(1)

	go func() {

		defer wg.Done()

		for i := 0; i < 5; i++ {

			atomic.StoreUint64(&counter, uint64(i+1))

			fmt.Printf("Writer: Counter set to %d\n", i+1)

			time.Sleep(time.Millisecond * 500)

		}

	}()

	// Reader goroutine
	wg.Add(1)

	go func() {

		defer wg.Done()

		for i := 0; i < 10; i++ {

			fmt.Printf("Reader: Current counter value is %d\n", atomic.LoadUint64(&counter))

			time.Sleep(time.Millisecond * 250)

		}

	}()

	wg.Wait()

}
