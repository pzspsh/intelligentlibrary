/*
@File   : jumpgame.go
@Author : pan
@Time   : 2023-05-23 10:35:15
*/
package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{3, 2, 2, 0, 4}))
}

// 跳跃游戏
func canJump(nums []int) bool {
	j := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= j {
			j = i
		}
	}
	return j <= 0
}
