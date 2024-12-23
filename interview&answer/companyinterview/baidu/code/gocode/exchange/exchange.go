/*
@File   : exchange.go
@Author : pan
@Time   : 2023-05-24 11:03:35
*/
package main

import "fmt"

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}

// 调整数组顺序使奇数位于偶数前面
func exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]%2 == 1 {
			i++
		} else if nums[j]%2 == 0 {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
