package main

func main() {
	root := TreeNode1{Val: 2}
	rootfirst := TreeNode1{Val: 3}
	rootSecond := TreeNode1{Val: 1}
	root.Left = &rootfirst
	root.Right = &rootSecond
	recoverTree1(&root)
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

// 恢复二叉搜索树
var prev, first, second *TreeNode1

func recoverTree1(root *TreeNode1) {
	prev, first, second = nil, nil, nil
	dfs1(root)
	first.Val, second.Val = second.Val, first.Val
}

func dfs1(root *TreeNode1) {
	if root == nil {
		return
	}
	dfs1(root.Left)
	if prev != nil && prev.Val > root.Val {
		second = root
		if first == nil {
			first = prev
		} else {
			return
		}
	}
	prev = root
	dfs1(root.Right)
}
