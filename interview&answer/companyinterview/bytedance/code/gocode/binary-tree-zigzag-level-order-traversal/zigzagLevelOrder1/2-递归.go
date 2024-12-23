package main

import "fmt"

func main() {
	root := TreeNode1{Val: 1}
	rootfirst := TreeNode1{Val: 2}
	rootSecond := TreeNode1{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(zigzagLevelOrder1(&root))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

var res [][]int

func zigzagLevelOrder1(root *TreeNode1) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	dfs1(root, 0)
	return res
}

func dfs1(root *TreeNode1, level int) {
	if root == nil {
		return
	}
	if level == len(res) {
		res = append(res, []int{})
	}
	if level%2 == 1 {
		arr := res[level]
		arr = append([]int{root.Val}, arr...)
		res[level] = arr
	} else {
		res[level] = append(res[level], root.Val)
	}
	dfs1(root.Left, level+1)
	dfs1(root.Right, level+1)
}
