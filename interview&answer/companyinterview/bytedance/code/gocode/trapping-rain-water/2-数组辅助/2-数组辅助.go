package main

import "fmt"

func main() {
	fmt.Println(trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap1(height []int) int {
	res := 0
	if len(height) == 0 {
		return 0
	}
	left := make([]int, len(height))
	right := make([]int, len(height))
	left[0] = height[0]
	right[len(right)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		left[i] = max1(height[i], left[i-1])
	}
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max1(height[i], right[i+1])
	}
	for i := 0; i < len(height); i++ {
		// 当前坐标形成的面积=(min(左边最高，右边最高)-当前高度) * 宽度(1,可省略)
		area := min1(left[i], right[i]) - height[i]
		res = res + area
	}
	return res
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min1(a, b int) int {
	if a > b {
		return b
	}
	return a
}
