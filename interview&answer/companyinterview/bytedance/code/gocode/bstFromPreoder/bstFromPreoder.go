/*
@File   : bstFromPreoder.go
@Author : pan
@Time   : 2023-07-04 23:29:42
*/
package main

import "fmt"

func main() {
	treenode := bstFromPreorder([]int{2, 4, 5, 6, 6, 4})
	fmt.Println(treenode)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历序列二叉搜索树
func bstFromPreorder(preorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}
	index := length
	for i := 1; i < length; i++ {
		if preorder[i] > preorder[0] {
			index = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:index]),
		Right: bstFromPreorder(preorder[index:]),
	}
}
