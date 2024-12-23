/*
@File   : treemaxdepth.go
@Author : pan
@Time   : 2023-05-23 10:10:32
*/
package main

import (
	"fmt"
)

func main() {
	root := TreeNode{}
	root.Val = 1

	left := TreeNode{}
	left.Val = 2

	right := TreeNode{}
	right.Val = 4

	//right1 := TreeNode{}
	//right1.Val = 2

	root.Left = &left
	left.Right = &right
	//right.Left = &right1
	fmt.Println(root.Left.Right)
	fmt.Println(maxDepth(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	//fmt.Println("aaaaa",left)
	//fmt.Println("bbbbb",right)
	//fmt.Println(max(left, right) + 1)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
