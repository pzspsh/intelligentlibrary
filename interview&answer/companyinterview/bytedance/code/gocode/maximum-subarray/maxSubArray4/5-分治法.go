package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray4([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray4([]int{1, 2}))
}

func maxSubArray4(nums []int) int {
	result := maxSubArr4(nums, 0, len(nums)-1)
	return result
}

func maxSubArr4(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2
	leftSum := maxSubArr4(nums, left, mid)        // 最大子序在左边
	rightSum := maxSubArr4(nums, mid+1, right)    // 最大子序在右边
	midSum := findMaxArr4(nums, left, mid, right) // 跨中心
	result := max(leftSum, rightSum)
	result = max(result, midSum)
	return result
}

func findMaxArr4(nums []int, left, mid, right int) int {
	leftSum := math.MinInt32
	sum := 0
	// 从右到左
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	rightSum := math.MinInt32
	sum = 0
	// 从左到右
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}
	return leftSum + rightSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
