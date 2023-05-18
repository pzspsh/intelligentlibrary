/*
@File   : subsets.go
@Author : pan
@Time   : 2023-05-18 14:13:09
*/
package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

// leetcode78_子集
var res3 [][]int

func subsets(nums []int) [][]int {
	res3 = make([][]int, 0)
	dfs(nums, make([]int, 0), 0)
	return res3
}

func dfs(nums []int, arr []int, level int) {
	if level >= len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res3 = append(res3, temp)
		return
	}
	dfs(nums, arr, level+1)
	dfs(nums, append(arr, nums[level]), level+1)
}
