package main

import (
	"fmt"
	"math"
)

func main() {
	root := TreeNode3{Val: 1}
	rootfirst := TreeNode3{Val: 2}
	rootSecond := TreeNode3{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(isValidBST3(&root))
}

type TreeNode3 struct {
	Val   int
	Left  *TreeNode3
	Right *TreeNode3
}

func isValidBST3(root *TreeNode3) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode3, 0)
	pre := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		if stack[last].Val <= pre {
			return false
		}
		pre = stack[last].Val
		root = stack[last].Right
		stack = stack[:last]
	}
	return true
}
