/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:07:34
*/
package main

import (
	"fmt"
)

func main() {
	var arr [3]int
	fmt.Println(arr)
	fmt.Println(cap(arr))
	var arr1 *[3]int
	fmt.Println(arr1)
	fmt.Println(cap(arr1))
	var s []string
	fmt.Println(s)
	fmt.Println(cap(s))
	s = make([]string, 1)
	s[0] = "go"
	fmt.Println(s)
	fmt.Println(cap(s))
	var c chan int
	fmt.Println(c)
	fmt.Println(cap(c))
}
