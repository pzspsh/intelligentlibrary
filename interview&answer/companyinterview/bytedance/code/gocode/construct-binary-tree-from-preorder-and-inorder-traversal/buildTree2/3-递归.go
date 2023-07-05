package main

import "fmt"

func main() {
	fmt.Println(buildTree2([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func buildTree2(preorder []int, inorder []int) *TreeNode2 {
	if len(preorder) == 0 {
		return nil
	}
	return helper(preorder, inorder)
}

func helper(preorder []int, inorder []int) *TreeNode2 {
	var root *TreeNode2
	for k := range inorder {
		if inorder[k] == preorder[0] {
			root = &TreeNode2{Val: preorder[0]}
			root.Left = helper(preorder[1:k+1], inorder[0:k])
			root.Right = helper(preorder[k+1:], inorder[k+1:])
		}
	}
	return root
}
