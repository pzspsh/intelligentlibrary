/*
@File   : getintersectionnode.go
@Author : pan
@Time   : 2023-05-19 14:11:53
*/
package main

import "fmt"

func main() {
	a := ListNode{Val: 10}
	b := ListNode{Val: 99}
	c := ListNode{Val: 100}
	d := ListNode{Val: 80}

	a.Next = &b
	d.Next = &b
	b.Next = &c
	fmt.Println(getIntersectionNode(&a, &d))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	A, B := headA, headB
	for A != B {
		if A != nil {
			A = A.Next
		} else {
			A = headB
		}
		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
	}
	return A
}
