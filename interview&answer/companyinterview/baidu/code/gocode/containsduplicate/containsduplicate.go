/*
@File   : containsduplicate.go
@Author : pan
@Time   : 2023-05-24 10:15:22
*/
package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(arr))
}

// 存在重复元素
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = 1
		}
	}
	return false
}
