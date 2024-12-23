package main

import "fmt"

func main() {
	fmt.Println(search1([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

func search1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
