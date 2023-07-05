package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring4("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring(" "))
}

func lengthOfLongestSubstring4(s string) int {
	arr := [256]int{}
	for i := range arr {
		arr[i] = -1
	}
	res, j := 0, -1
	for i := 0; i < len(s); i++ {
		if arr[s[i]] > j { // 出现重复了，更新下标
			j = arr[s[i]]
		} else {
			res = max(res, i-j) // 没有重复，更新长度
		}
		arr[s[i]] = i
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
