package main

import "fmt"

func main() {
	sum := twoSum([]int{3, 5, 14, 12}, 26)
	fmt.Println(sum)
}

// 哈希遍历之和
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
