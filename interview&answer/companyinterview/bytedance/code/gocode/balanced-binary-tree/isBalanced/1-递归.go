package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 2

	root.Left = &left
	left.Right = &right
	res := isBalanced(&root)
	fmt.Println(res)
}

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	_, isBalanced := recur(root)
	return isBalanced
}

func recur(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	leftDepth, leftIsBalanced := recur(root.Left)
	if !leftIsBalanced {
		return 0, false
	}
	rightDepth, rightIsBalanced := recur(root.Right)
	if !rightIsBalanced {
		return 0, false
	}

	if -1 <= leftDepth-rightDepth &&
		leftDepth-rightDepth <= 1 {
		return max(leftDepth, rightDepth) + 1, true
	}
	return 0, false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
