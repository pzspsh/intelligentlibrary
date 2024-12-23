/*
@File   : hammingweight.go
@Author : pan
@Time   : 2023-05-23 10:19:16
*/
package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(3))
	fmt.Println(hammingWeight(4))
}

// 位1的个数
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
