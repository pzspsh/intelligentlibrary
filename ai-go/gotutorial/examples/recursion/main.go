/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:38:07
*/
package main

import (
	"fmt"
)

// 下面例子展示了函数递归的用法。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}
