/*
@File   : swapnumbers.go
@Author : pan
@Time   : 2023-05-24 10:43:33
*/
package main

import "fmt"

func main() {
	fmt.Println(swapNumbers([]int{1, 2}))
}

// 交换数字
func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] + numbers[1]
	numbers[1] = numbers[0] - numbers[1]
	numbers[0] = numbers[0] - numbers[1]
	return numbers
}
