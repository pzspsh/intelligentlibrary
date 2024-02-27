/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:06:52
*/
package main

import (
	"fmt"
)

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)
	fmt.Println(len(arr))
	var arr1 *[3]int
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	var s []int
	fmt.Println(len(s))
	s = []int{1, 2, 3}
	fmt.Println(len(s))
	var m map[int]string
	fmt.Println(len(m))
	m = make(map[int]string)
	m[0] = "hello"
	fmt.Println(len(m))
	str := "frank"
	fmt.Println(len(str))
	var c chan int
	fmt.Println(c)
	fmt.Println(len(c))
	c = make(chan int)
	fmt.Println(len(c))
}
