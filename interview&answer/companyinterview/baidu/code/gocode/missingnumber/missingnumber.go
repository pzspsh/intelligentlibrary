/*
@File   : missingnumber.go
@Author : pan
@Time   : 2023-05-19 13:19:35
*/
package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	fmt.Println(missingNumber([]int{3, 0, 1}))
}

// 缺失数字
func missingNumber(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum = sum - nums[i]
	}
	return sum
}
