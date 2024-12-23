package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix1([]string{"flower", "flow", "flight"}))
	// fmt.Println(longestCommonPrefix([]string{"a"}))
}

// 最长公共前缀
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	length := 0

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || char != strs[j][i] {
				return strs[0][:length]
			}
		}
		length++
	}
	return strs[0][:length]
}
