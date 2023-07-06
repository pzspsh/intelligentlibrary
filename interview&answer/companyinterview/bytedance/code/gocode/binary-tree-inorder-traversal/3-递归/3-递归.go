package main

import "fmt"

func main() {
	root := TreeNode2{Val: 1}
	rootfirst := TreeNode2{Val: 2}
	rootSecond := TreeNode2{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(inorderTraversal2(&root))
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

var res []int

func inorderTraversal2(root *TreeNode2) []int {
	res = make([]int, 0)
	dfs2(root)
	return res
}

func dfs2(root *TreeNode2) {
	if root != nil {
		dfs2(root.Left)
		res = append(res, root.Val)
		dfs2(root.Right)
	}
}
