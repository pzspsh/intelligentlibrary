package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea1([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea1(heights []int) int {
	n := len(heights)
	res := 0
	for i := 0; i < n; i++ {
		height := heights[i]
		left, right := i, i
		for left > 0 && heights[left-1] >= height {
			left--
		}
		for right < n-1 && heights[right+1] >= height {
			right++
		}
		width := right - left + 1
		res = max1(res, width*height)
	}
	return res
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
