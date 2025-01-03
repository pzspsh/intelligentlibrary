/*
@File   : majorityelement.go
@Author : pan
@Time   : 2023-05-25 10:19:19
*/
package main

import "fmt"

func main() {
	nums := []int{0, 3, 2, 2, 1, 1, 1, 2, 2, 1, 1}
	fmt.Println(majorityElement(nums))
}

// 多数元素
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
