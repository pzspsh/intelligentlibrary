/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:19:56
*/
package main

import (
	"fmt"

	"sync"

	"sync/atomic"

	"unsafe"
)

type sharedStruct struct {
	value int
}

func main() {

	s := &sharedStruct{value: 100}

	sharedPointer := unsafe.Pointer(s)

	var wg sync.WaitGroup

	read := func(id int) {

		defer wg.Done()

		p := atomic.LoadPointer(&sharedPointer)

		sharedObj := (*sharedStruct)(p)

		fmt.Printf("Goroutine %d - Read value: %d\n", id, sharedObj.value)

	}

	wg.Add(5)

	for i := 1; i <= 5; i++ {

		go read(i)

	}

	wg.Wait()

}
