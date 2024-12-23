/*
@File   : maxsubarray.go
@Author : pan
@Time   : 2023-05-18 14:06:07
*/
package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{1, 2}))
}

// 最大子序和
func maxSubArray(nums []int) int {
	result := nums[0]
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > result {
			result = sum
		}
	}
	return result
}
