package main

import "fmt"

func main() {
	// sum := twoSum([]int{2, 7, 11, 15, 24}, 26) // [0 4]
	sum := twoSum([]int{2, 7, 11, 15}, 26) // [2 3]
	fmt.Println(sum)
}

// 两数之和 暴力法
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
