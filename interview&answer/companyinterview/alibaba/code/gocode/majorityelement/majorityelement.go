/*
@File   : majorityelement.go
@Author : pan
@Time   : 2023-05-19 14:08:32
*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

// 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	result, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
		} else if result == nums[i] {
			count++
		} else {
			count--
		}
	}
	return result
}
