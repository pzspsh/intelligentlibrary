/*
@File   : hascycle.go
@Author : pan
@Time   : 2023-05-24 11:04:46
*/
package main

import "fmt"

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &second

	fmt.Println(hasCycle(&first))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head.Next
	for fast != nil && head != nil && fast.Next != nil {
		if fast == head {
			return true
		}
		fast = fast.Next.Next
		head = head.Next
	}
	return false
}
