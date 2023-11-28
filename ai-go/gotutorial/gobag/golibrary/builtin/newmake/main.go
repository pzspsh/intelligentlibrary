/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:02:00
*/
package main

import (
	"fmt"
)

func main() {
	var p *[]int = new([]int)     // allocates slice structure; *p == ni; rarely useful
	var v []int = make([]int, 10) // the slice v now refers to a new array of 100 ints

	fmt.Printf("p: %v\n", p)
	fmt.Printf("v: %v\n", v)

	// Unnecessarily complex: 这种做法实在是很蛋疼

	var p1 *[]int = new([]int)
	*p1 = make([]int, 5, 10)
	// Idiomatic: 习惯的做法
	v1 := make([]int, 10)

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("v1: %v\n", v1)
}
