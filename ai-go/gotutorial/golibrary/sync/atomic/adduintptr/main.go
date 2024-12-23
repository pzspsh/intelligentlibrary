/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 17:18:37
*/
package main

import (
	"fmt"

	"sync"

	"sync/atomic"
)

const bufferSize = 100

type Buffer struct {
	data [bufferSize]byte

	offset uintptr
}

func (b *Buffer) Write(p []byte) (n int, err error) {

	n = copy(b.data[b.offset:], p)

	atomic.AddUintptr(&b.offset, uintptr(n))

	return

}

func main() {

	var buf Buffer

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {

		defer wg.Done()

		data := []byte("Hello, ")

		n, _ := buf.Write(data)

		fmt.Printf("Goroutine 1 wrote %d bytes\n", n)

	}()

	go func() {

		defer wg.Done()

		data := []byte("World!")

		n, _ := buf.Write(data)

		fmt.Printf("Goroutine 2 wrote %d bytes\n", n)

	}()

	wg.Wait()

	fmt.Printf("Buffer content: %s\n", buf.data[:buf.offset])

}
