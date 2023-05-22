/*
@File   : trianglenumber.go
@Author : pan
@Time   : 2023-05-22 09:57:16
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4}))
}

// 有效三角形的个数
func triangleNumber(nums []int) int {
	res := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] > nums[k] &&
					nums[i]+nums[k] > nums[j] &&
					nums[j]+nums[k] > nums[i] {
					res++
				}
			}
		}
	}
	return res
}
