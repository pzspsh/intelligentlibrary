package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray1([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray1([]int{1, 2}))
}

func maxSubArray1(nums []int) int {
	result := math.MinInt32

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > result {
				result = sum
			}
		}
	}
	return result
}
