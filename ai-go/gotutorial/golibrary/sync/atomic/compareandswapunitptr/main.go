package main

import (
	"fmt"

	"sync"

	"sync/atomic"
)

func main() {

	var sharedValue uintptr

	var wg sync.WaitGroup

	update := func(id int) {

		defer wg.Done()

		for i := 0; i < 10; i++ {

			for {

				old := atomic.LoadUintptr(&sharedValue)

				new := old + uintptr(id)

				if atomic.CompareAndSwapUintptr(&sharedValue, old, new) {

					break

				}

			}

		}

	}

	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go update(i)

	}

	wg.Wait()

	fmt.Println("Final shared value:", sharedValue)

}
