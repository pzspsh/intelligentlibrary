/*
@File   : singlenumber.go
@Author : pan
@Time   : 2023-05-19 12:58:54
*/
package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}

// 只出现一次的数字
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
