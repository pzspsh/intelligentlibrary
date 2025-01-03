package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray2([]int{1, 2}))
}

// dp[i]=max(dp[i-1]+nums[i],nums[i])
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
