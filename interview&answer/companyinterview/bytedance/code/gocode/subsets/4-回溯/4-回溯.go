package main

import "fmt"

func main() {
	fmt.Println(subsets3([]int{1, 2, 3}))
}

// 子集
var res3 [][]int

func subsets3(nums []int) [][]int {
	res3 = make([][]int, 0)
	dfs3(nums, make([]int, 0), 0)
	return res3
}

func dfs3(nums []int, arr []int, level int) {
	if level >= len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res3 = append(res3, temp)
		return
	}
	dfs3(nums, arr, level+1)
	dfs3(nums, append(arr, nums[level]), level+1)
}
