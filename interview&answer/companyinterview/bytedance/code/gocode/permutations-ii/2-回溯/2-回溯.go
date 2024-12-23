package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permuteUnique1([]int{1, 1, 2}))
}

var res1 [][]int

func permuteUnique1(nums []int) [][]int {
	res1 = make([][]int, 0)
	sort.Ints(nums)
	dfs1(nums, 0)
	return res1
}

func dfs1(nums []int, index int) {
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		res1 = append(res1, temp)
		return
	}
	m := make(map[int]int)
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		m[nums[i]] = 1
		nums[i], nums[index] = nums[index], nums[i]
		dfs1(nums, index+1)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
