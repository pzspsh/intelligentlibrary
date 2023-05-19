/*
@File   : mirrortree.go
@Author : pan
@Time   : 2023-05-19 14:07:40
*/
package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	mirrorTree(&first)
	fmt.Println(first.Val)
	fmt.Println(first.Left.Val)
	fmt.Println(first.Right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	root.Left, root.Right = mirrorTree(root.Right), mirrorTree(root.Left)
	return root
}
