/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:49:45
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

var data *string

// get data atomically
func Data() string {
	p := (*string)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
	))
	if p == nil {
		return ""
	} else {
		return *p
	}
}

// set data atomically
func SetData(d string) {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&data)),
		unsafe.Pointer(&d),
	)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(200)

	for range [100]struct{}{} {
		go func() {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)

			log.Println(Data())
			wg.Done()
		}()
	}

	for i := range [100]struct{}{} {
		go func(i int) {
			time.Sleep(time.Second * time.Duration(rand.Intn(1000)) / 1000)
			s := fmt.Sprint("#", i)
			log.Println("====", s)

			SetData(s)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("final data = ", *data)
}
