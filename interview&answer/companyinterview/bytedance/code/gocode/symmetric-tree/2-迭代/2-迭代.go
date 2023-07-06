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
	right.Val = 2

	root.Left = &left
	root.Right = &right

	fmt.Println(isSymmetric1(&root))
}

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func isSymmetric1(root *TreeNode1) bool {
	leftQ := make([]*TreeNode1, 0)
	rightQ := make([]*TreeNode1, 0)
	leftQ = append(leftQ, root)
	rightQ = append(rightQ, root)

	for len(leftQ) != 0 && len(rightQ) != 0 {
		leftCur, rightCur := leftQ[0], rightQ[0]
		leftQ, rightQ = leftQ[1:], rightQ[1:]

		if leftCur == nil && rightCur == nil {
			continue
		} else if leftCur != nil && rightCur != nil && leftCur.Val == rightCur.Val {
			leftQ = append(leftQ, leftCur.Left, leftCur.Right)
			rightQ = append(rightQ, rightCur.Right, rightCur.Left)
		} else {
			return false
		}
	}

	if len(leftQ) == 0 && len(rightQ) == 0 {
		return true
	} else {
		return false
	}
}
