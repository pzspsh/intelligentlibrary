/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:06:02
*/
package main

import (
	"fmt"
)

func main() {
	var m map[int]string
	fmt.Println(m)
	delete(m, 0)
	fmt.Println(m)
	m1 := make(map[int]string)
	fmt.Println(m1)
	delete(m1, 0)
	fmt.Println(m1)
	m2 := make(map[int]string, 2)
	m2[0] = "hello"
	m2[1] = "world"
	fmt.Println(m2)
	delete(m2, 0)
	fmt.Println(m2)
}
