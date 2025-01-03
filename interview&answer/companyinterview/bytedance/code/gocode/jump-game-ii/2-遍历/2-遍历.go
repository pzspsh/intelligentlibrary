package main

import "fmt"

func main() {
	fmt.Println(jump1([]int{2, 3, 1, 1, 4}))
}

// 跳跃游戏II
func jump1(nums []int) int {
	res := 0
	end := 0
	maxValue := 0
	for i := 0; i < len(nums)-1; i++ {
		maxValue = max(maxValue, i+nums[i])
		if i == end {
			end = maxValue
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
