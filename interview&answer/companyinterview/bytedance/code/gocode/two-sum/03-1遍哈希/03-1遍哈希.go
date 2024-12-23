package main

import "fmt"

func main() {
	fmt.Println(twoSum3([]int{2, 7, 11, 15}, 26))
}

func twoSum3(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := m[target-b]; ok {
			return []int{j, i}
		}
		m[b] = i
	}
	return nil
}
