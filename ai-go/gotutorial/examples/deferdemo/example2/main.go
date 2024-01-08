/*
@File   : main.go
@Author : pan
@Time   : 2024-01-08 12:20:58
*/
package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		wg.Done()
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		wg.Done()
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
