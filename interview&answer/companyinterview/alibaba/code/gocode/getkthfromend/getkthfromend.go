/*
@File   : getkthfromend.go
@Author : pan
@Time   : 2023-05-19 14:05:14
*/
package main

import "fmt"

func main() {
	first := ListNode{Val: 1}
	second := ListNode{Val: 2}
	third := ListNode{Val: 3}
	first.Next = &second
	second.Next = &third
	fmt.Println(getKthFromEnd(&first, 2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	for k > 0 && head != nil {
		fast = fast.Next
		k--
	}
	if k > 0 {
		return nil
	}
	slow := head
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
