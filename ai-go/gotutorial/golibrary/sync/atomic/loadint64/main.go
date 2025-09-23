package main

import (
	"fmt"

	"sync"

	"sync/atomic"

	"time"
)

func main() {

	var sharedValue int64 = 100

	var wg sync.WaitGroup

	read := func(id int) {

		defer wg.Done()

		time.Sleep(time.Duration(id) * 10 * time.Millisecond)

		val := atomic.LoadInt64(&sharedValue)

		fmt.Printf("Goroutine %d - Read value: %d\n", id, val)

	}

	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go read(i)

	}

	wg.Wait()

}
