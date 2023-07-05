package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses2("((())"))
}

func longestValidParentheses2(s string) int {
	res := 0
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max2(res, 2*left)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			res = max2(res, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return res
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
