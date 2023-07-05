package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix2([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix2([]string{"a"}))
	fmt.Println(longestCommonPrefix2([]string{"abc", "acb", "bac"}))
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)
	first := strs[0]
	last := strs[len(strs)-1]
	i := 0
	length := len(first)
	if len(last) < length {
		length = len(last)
	}
	for i < length {
		if first[i] != last[i] {
			return first[:i]
		}
		i++
	}
	return first[:i]
}
