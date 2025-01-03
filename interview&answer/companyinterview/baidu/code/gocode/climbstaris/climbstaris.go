/*
@File   : climbstaris.go
@Author : pan
@Time   : 2023-05-18 14:09:18
*/
package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

// 爬楼梯
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	first := 1
	second := 2
	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}
	return second
}
