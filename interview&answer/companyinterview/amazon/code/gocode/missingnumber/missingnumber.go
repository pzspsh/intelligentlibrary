/*
@File   : missingnumber.go
@Author : pan
@Time   : 2023-05-22 09:40:03
*/
package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
}

// 0～n-1中缺失的数字
func missingNumber(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] != mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
