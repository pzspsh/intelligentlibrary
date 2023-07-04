/*
@File   : maxProduct.go
@Author : pan
@Time   : 2023-07-04 23:25:41
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

// 乘积最大子数组
func maxProduct(nums []int) int {
	res := math.MinInt64
	for i := 0; i < len(nums); i++ {
		temp := 1
		for j := i; j < len(nums); j++ {
			temp = temp * nums[j]
			if temp > res {
				res = temp
			}
		}
	}
	return res
}
