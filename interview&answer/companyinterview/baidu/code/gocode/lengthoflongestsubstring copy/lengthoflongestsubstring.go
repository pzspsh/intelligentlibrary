/*
@File   : lengthoflongestsubstring.go
@Author : pan
@Time   : 2023-05-19 14:09:37
*/
package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcabcbb"))
	fmt.Println(lengthOfLongestSubstring("aabaab!bb"))
	fmt.Println(lengthOfLongestSubstring("au"))
}

// 最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	res := 1
	arr := make(map[byte]int)
	curLength := 0
	for i := 0; i < len(s); i++ {
		if preIndex, ok := arr[s[i]]; !ok || i-preIndex > curLength {
			curLength++
		} else {
			if curLength > res {
				res = curLength
			}
			curLength = i - preIndex
		}
		arr[s[i]] = i
	}
	if curLength > res {
		res = curLength
	}
	return res
}
