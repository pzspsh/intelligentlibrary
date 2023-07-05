package main

import (
	"fmt"
)

func main() {
	root := TreeNode2{Val: 1}
	rootfirst := TreeNode2{Val: 2}
	rootSecond := TreeNode2{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(isValidBST2(&root))
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func isValidBST2(root *TreeNode2) bool {
	if root == nil {
		return true
	}
	stack := make([]*TreeNode2, 0)
	res := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		res = append(res, stack[last].Val)
		root = stack[last].Right
		stack = stack[:last]
	}
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}
