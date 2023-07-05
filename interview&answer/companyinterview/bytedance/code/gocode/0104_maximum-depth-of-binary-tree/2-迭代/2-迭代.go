package main

import (
	"fmt"
)

func main() {
	root := TreeNode1{}
	root.Val = 1

	left := TreeNode1{}
	left.Val = 2

	right := TreeNode1{}
	right.Val = 3

	root.Left = &left
	left.Right = &right

	fmt.Println(maxDepth1(&root))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func maxDepth1(root *TreeNode1) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode1, 0)
	queue = append(queue, root)
	depth := 0

	for len(queue) > 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
