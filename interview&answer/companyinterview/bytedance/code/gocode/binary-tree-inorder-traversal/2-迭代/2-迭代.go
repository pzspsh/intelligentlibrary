package main

import "fmt"

func main() {
	root := TreeNode1{Val: 1}
	rootfirst := TreeNode1{Val: 2}
	rootSecond := TreeNode1{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(inorderTraversal1(&root))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

// leetcode94_二叉树的中序遍历
func inorderTraversal1(root *TreeNode1) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode1, 0)
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
	return res
}
