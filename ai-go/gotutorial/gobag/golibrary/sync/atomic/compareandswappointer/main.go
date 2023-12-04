/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:21:49
*/
package main

import (
	"fmt"

	"sync"

	"sync/atomic"

	"unsafe"
)

type Data struct {
	value int
}

func main() {

	var dataPtr unsafe.Pointer

	var wg sync.WaitGroup

	data := &Data{value: 10}

	dataPtr = unsafe.Pointer(data)

	wg.Add(2)

	go func() {

		defer wg.Done()

		newData := &Data{value: 20}

		swapped := atomic.CompareAndSwapPointer(&dataPtr, unsafe.Pointer(data), unsafe.Pointer(newData))

		fmt.Println("Goroutine 1 - Swapped:", swapped)

	}()

	go func() {

		defer wg.Done()

		newData := &Data{value: 30}

		swapped := atomic.CompareAndSwapPointer(&dataPtr, unsafe.Pointer(data), unsafe.Pointer(newData))

		fmt.Println("Goroutine 2 - Swapped:", swapped)

	}()

	wg.Wait()

	finalData := (*Data)(dataPtr)

	fmt.Println("Final value:", finalData.value)

}
