package main

import "fmt"

func main() {
	fmt.Println(permute2([]int{1, 2, 3}))
}

var res2 [][]int

func permute2(nums []int) [][]int {
	res2 = make([][]int, 0)
	arr := make([]int, len(nums))
	dfs2(nums, 0, arr)
	return res2
}

func dfs2(nums []int, index int, arr []int) {
	if index == len(nums) {
		temp := make([]int, len(arr))
		copy(temp, arr)
		res2 = append(res2, temp)
		return
	}
	for i := index; i < len(nums); i++ {
		arr[index] = nums[i]
		nums[i], nums[index] = nums[index], nums[i]
		dfs2(nums, index+1, arr)
		nums[i], nums[index] = nums[index], nums[i]
	}
}
