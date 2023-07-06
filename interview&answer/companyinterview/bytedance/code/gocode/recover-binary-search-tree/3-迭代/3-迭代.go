package main

func main() {
	root := TreeNode2{Val: 2}
	rootfirst := TreeNode2{Val: 3}
	rootSecond := TreeNode2{Val: 1}
	root.Left = &rootfirst
	root.Right = &rootSecond
	recoverTree2(&root)
}

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func recoverTree2(root *TreeNode2) {
	var prev, first, second *TreeNode2
	stack := make([]*TreeNode2, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && root.Val < prev.Val {
			second = root
			if first == nil {
				first = prev
			} else {
				break
			}
		}
		prev = root
		root = root.Right
	}
	first.Val, second.Val = second.Val, first.Val
}
