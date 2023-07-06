/*
@File   : isvalidbst.go
@Author : pan
@Time   : 2023-05-24 11:54:40
*/
package main

import (
	"fmt"
)

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(isValidBST(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 验证二叉搜索树
var res []int

func isValidBST(root *TreeNode) bool {
	res = make([]int, 0)
	dfs(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}
}
