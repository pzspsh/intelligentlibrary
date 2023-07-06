package main

import "fmt"

func main() {
	fmt.Println(largestRectangleArea3([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea3([]int{1}))
}

func largestRectangleArea3(heights []int) int {
	n := len(heights)
	res := 0
	left := make([]int, n)
	right := make([]int, n)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		res = max3(res, heights[i]*(right[i]-left[i]-1))
	}
	return res
}

func max3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
