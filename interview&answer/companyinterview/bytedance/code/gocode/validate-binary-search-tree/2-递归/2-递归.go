package main

import (
	"fmt"
)

func main() {
	root := TreeNode1{Val: 1}
	rootfirst := TreeNode1{Val: 2}
	rootSecond := TreeNode1{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(isValidBST1(&root))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

// leetcode98_验证二叉搜索树
var res []int

func isValidBST1(root *TreeNode1) bool {
	res = make([]int, 0)
	dfs1(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] {
			return false
		}
	}
	return true
}

func dfs1(root *TreeNode1) {
	if root != nil {
		dfs1(root.Left)
		res = append(res, root.Val)
		dfs1(root.Right)
	}
}
