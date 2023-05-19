/*
@File   : minarray.go
@Author : pan
@Time   : 2023-05-19 14:04:23
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
}

// 旋转数组的最小数字
func minArray(numbers []int) int {
	for i := 1; i < len(numbers); i++ {
		if numbers[i] < numbers[i-1] {
			return numbers[i]
		}
	}
	return numbers[0]
}
