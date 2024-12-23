/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:22:57
*/
package main

import (
	"fmt"

	"sync"

	"sync/atomic"
)

func main() {

	var counter uint64

	var wg sync.WaitGroup

	increment := func() {

		defer wg.Done()

		for i := 0; i < 1000; i++ {

			for {

				old := atomic.LoadUint64(&counter)

				new := old + 1

				if atomic.CompareAndSwapUint64(&counter, old, new) {

					break

				}

			}

		}

	}

	wg.Add(5)

	for i := 0; i < 5; i++ {

		go increment()

	}

	wg.Wait()

	fmt.Println("Final counter value:", counter)

}
