/*
@File   : zigzaglevleorder.go
@Author : pan
@Time   : 2023-05-24 11:22:35
*/
package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond
	fmt.Println(zigzagLevelOrder(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的锯齿形层次遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		length := len(list)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := list[i]
			temp = append(temp, node.Val)
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
		}
		if len(res)%2 == 1 {
			for i := 0; i < len(temp)/2; i++ {
				temp[i], temp[len(temp)-1-i] = temp[len(temp)-1-i], temp[i]
			}
		}
		res = append(res, temp)
		list = list[length:]
	}
	return res
}
