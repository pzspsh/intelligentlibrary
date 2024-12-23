package main

import "fmt"

func main() {
	fmt.Println(buildTree2([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func buildTree2(inorder []int, postorder []int) *TreeNode2 {
	if postorder == nil || len(postorder) == 0 {
		return nil
	}
	last := len(postorder) - 1
	root := &TreeNode2{
		Val: postorder[last],
	}
	length := len(postorder)
	stack := make([]*TreeNode2, 0)
	stack = append(stack, root)
	index := last
	for i := length - 2; i >= 0; i-- {
		value := postorder[i]
		node := stack[len(stack)-1]
		if node.Val != inorder[index] {
			node.Right = &TreeNode2{Val: value}
			stack = append(stack, node.Right)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				index--
			}
			node.Left = &TreeNode2{Val: value}
			stack = append(stack, node.Left)
		}
	}
	return root
}
