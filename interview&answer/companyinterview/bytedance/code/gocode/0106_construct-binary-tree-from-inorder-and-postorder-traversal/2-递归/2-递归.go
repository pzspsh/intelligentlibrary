package main

import "fmt"

func main() {
	fmt.Println(buildTree1([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func buildTree1(inorder []int, postorder []int) *TreeNode1 {
	if len(postorder) == 0 {
		return nil
	}
	return helper(inorder, postorder)
}

func helper(inorder []int, postorder []int) *TreeNode1 {
	var root *TreeNode1
	last := len(postorder) - 1
	for k := range inorder {
		if inorder[k] == postorder[last] {
			root = &TreeNode1{Val: postorder[last]}
			root.Left = helper(inorder[0:k], postorder[0:k])
			root.Right = helper(inorder[k+1:], postorder[k:last])
		}
	}
	return root
}
