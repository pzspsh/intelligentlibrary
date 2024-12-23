package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum1([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

var res1 [][]int

func combinationSum1(candidates []int, target int) [][]int {
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
		if i != index && candidates[i] == candidates[i-1] {
			continue
		}
		if target < 0 {
			return
		}
		arr = append(arr, candidates[i])
		dfs1(candidates, target-candidates[i], arr, i+1)
		arr = arr[:len(arr)-1]
	}
}
