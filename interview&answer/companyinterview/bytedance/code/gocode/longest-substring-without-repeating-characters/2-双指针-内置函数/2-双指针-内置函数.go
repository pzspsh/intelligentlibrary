package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring1(str))
}

func lengthOfLongestSubstring1(s string) int {
	max, j := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[j:i], string(s[i]))
		if index == -1 {
			continue
		}
		if i-j > max {
			max = i - j
		}
		j = j + index + 1
	}
	if len(s)-j > max {
		max = len(s) - j
	}
	return max
}
