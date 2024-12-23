/*
@File   : reversewords.go
@Author : pan
@Time   : 2023-05-19 13:51:41
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}

// 反转字符串中的单词
func reverseWords(s string) string {
	arr := []byte(s)
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == ' ' {
			reverse(arr, j, i-1)
			j = i + 1
		}
	}
	reverse(arr, j, len(arr)-1)
	return string(arr)
}

func reverse(arr []byte, i, j int) []byte {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}