package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive3([]int{1, 2, 0}))
}

func firstMissingPositive3(nums []int) int {
	n := len(nums)
	for i := 1; i <= n; i++ {
		flag := false
		for j := 0; j < n; j++ {
			if i == nums[j] {
				flag = true
				break
			}
		}
		if flag == false {
			return i
		}
	}
	return n + 1
}
