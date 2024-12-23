/*
@File   : maxproduct.go
@Author : pan
@Time   : 2023-05-24 11:45:25
*/
package main

import "fmt"

func main() {
	root := &TreeNode{}
	res := maxProduct(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分裂二叉树的最大乘积
var sum int
var res int

func maxProduct(root *TreeNode) int {
	sum = 0
	res = 0
	dfsSum(root)
	dfs(root)
	return res % 1000000007
}

func dfsSum(root *TreeNode) {
	if root == nil {
		return
	}
	sum = sum + root.Val
	dfsSum(root.Left)
	dfsSum(root.Right)
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	res = max(res, left*(sum-left))
	res = max(res, right*(sum-right))
	return left + right + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
