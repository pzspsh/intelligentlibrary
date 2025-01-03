/*
@File   : towsum.go
@Author : pan
@Time   : 2023-05-18 13:49:22
*/
package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}

// 两数之和
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for k, v := range nums {
		m[v] = k
	}

	for i := 0; i < len(nums); i++ {
		b := target - nums[i]
		if num, ok := m[b]; ok && num != i {
			return []int{i, m[b]}
		}
	}
	return []int{}
}
