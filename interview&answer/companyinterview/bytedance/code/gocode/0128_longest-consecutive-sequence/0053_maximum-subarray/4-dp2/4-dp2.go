package main

import "fmt"

func main() {
	fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray3([]int{1, 2}))
}

func maxSubArray3(nums []int) int {
	dp := nums[0]
	result := dp

	for i := 1; i < len(nums); i++ {
		if dp+nums[i] > nums[i] {
			dp = dp + nums[i]
		} else {
			dp = nums[i]
		}

		if dp > result {
			result = dp
		}
	}
	return result
}
