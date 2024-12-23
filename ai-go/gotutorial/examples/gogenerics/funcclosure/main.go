/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 13:53:18
*/
package main

import (
	"fmt"
	"strconv"
)

func Filter[T any](src []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, t := range src {
		if f(t) {
			res = append(res, t)
		}
	}
	return res
}

func Map[S, T any](src []S, f func(S) T) []T {
	res := make([]T, 0)
	for _, s := range src {
		t := f(s)
		res = append(res, t)
	}
	return res
}

func Reduce[T any](src []T, f func(T, T) T) T {
	if len(src) == 1 {
		return src[0]
	}

	return f(src[0], Reduce(src[1:], f))
}

// 测试函数闭包
func main() {
	// filter test
	filterTest := Filter[int]([]int{1, 2, 3, 4, 5}, func(i int) bool {
		if i > 3 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(filterTest)

	// map test
	mapTest := Map[int, string]([]int{1, 2, 3}, func(i int) string {
		return "str" + strconv.Itoa(i)
	})
	fmt.Println(mapTest)

	// reduce test
	reduceTest := Reduce([]int{1, 2, 3}, func(a int, b int) int {
		return a + b
	})
	fmt.Println(reduceTest)
}
