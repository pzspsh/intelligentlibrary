package main

import (
	"fmt"
	"math"
)

func main() {
	root := TreeNode4{Val: 0}
	//rootfirst := TreeNode{Val: 2}
	//rootSecond := TreeNode{Val: 3}
	//root.Left = &rootfirst
	//root.Right = &rootSecond
	fmt.Println(isValidBST4(&root))
}

type TreeNode4 struct {
	Val   int
	Left  *TreeNode4
	Right *TreeNode4
}

var pre int

func isValidBST4(root *TreeNode4) bool {
	pre = math.MinInt64
	return dfs4(root)
}

func dfs4(root *TreeNode4) bool {
	if root == nil {
		return true
	}
	if dfs4(root.Left) == false {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return dfs4(root.Right)
}
