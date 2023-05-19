/*
@File   : fib.go
@Author : pan
@Time   : 2023-05-19 13:46:16
*/
package main

import "fmt"

func main() {
	fmt.Println(fib(2))
	fmt.Println(fib(3))
}

// 斐波那契数
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	n1, n2 := 0, 1
	for i := 2; i <= N; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}
