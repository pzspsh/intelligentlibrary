/*
@File   : uniquepaths.go
@Author : pan
@Time   : 2023-05-23 10:03:37
*/
package main

import "fmt"

func main() {
	fmt.Println(uniquePaths2(3, 2))
}

// 不同路径
// 求C((m+n-2),(m-1))=> (n+m-2)!/(m-1)!(n-1)!
func uniquePaths2(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if m > n {
		m, n = n, m
	}
	a := 1
	for i := 1; i <= m-1; i++ {
		a = a * i
	}
	b := 1
	for i := n; i <= m+n-2; i++ {
		b = b * i
	}
	return b / a
}
