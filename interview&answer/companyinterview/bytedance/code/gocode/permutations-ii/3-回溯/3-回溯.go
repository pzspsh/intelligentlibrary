package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique2([]int{1, 1, 2}))
	fmt.Println(permuteUnique2([]int{1, 2, 3}))
	fmt.Println(permuteUnique2([]int{0, 1, 0, 0, 9}))
}

// 全排列II
var res2 [][]int

func permuteUnique2(nums []int) [][]int {
	res2 = make([][]int, 0)
	sort.Ints(nums)
	dfs2(nums, make([]int, 0))
	return res2
}

func dfs2(nums []int, arr []int) {
	if len(nums) == 0 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res2 = append(res2, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		tempArr := make([]int, len(nums))
		copy(tempArr, nums)
		arr = append(arr, nums[i])
		dfs2(append(tempArr[:i], tempArr[i+1:]...), arr)
		arr = arr[:len(arr)-1]
	}
}
