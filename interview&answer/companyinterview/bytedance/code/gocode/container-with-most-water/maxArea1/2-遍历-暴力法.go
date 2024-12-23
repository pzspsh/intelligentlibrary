package main

import "fmt"

func main() {
	fmt.Println(maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea1(height []int) int {
	res := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min1(height[i], height[j])
			if area > res {
				res = area
			}
		}
	}
	return res
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
