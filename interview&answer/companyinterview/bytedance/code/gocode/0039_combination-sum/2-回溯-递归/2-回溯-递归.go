package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

// leetcode39_组合总和
var res1 [][]int

func combinationSum(candidates []int, target int) [][]int {
	res1 = make([][]int, 0)
	sort.Ints(candidates)
	dfs1(candidates, target, []int{}, 0)
	return res1
}

func dfs1(candidates []int, target int, arr []int, index int) {
	if target == 0 {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res1 = append(res1, temp)
		return
	}
	for i := index; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		dfs1(candidates, target-candidates[i], append(arr, candidates[i]), i)
	}
}
