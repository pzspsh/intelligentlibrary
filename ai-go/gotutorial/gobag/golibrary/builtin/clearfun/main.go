/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:24:56
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("len=%d\t s=%+v\n", len(s), s) // len=3      s=[1 2 3]
	clear(s)
	fmt.Printf("len=%d\t s=%+v\n", len(s), s) // len=3      s=[0 0 0]

	m := map[string]int{"go": 100, "php": 80}
	fmt.Printf("len=%d\tm=%+v\n", len(m), m) // len=2   m=map[go:100 php:80]
	clear(m)
	fmt.Printf("len=%d\tm=%+v\n", len(m), m) // len=0   m=map[]
}
