/*
@File   : sortlist.go
@Author : pan
@Time   : 2023-05-24 11:08:18
*/
package main

import "fmt"

func main() {
	head := &ListNode{}
	res := sortList(head)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

func quickSort(head, end *ListNode) {
	if head == end || head.Next == end {
		return
	}
	temp := head.Val
	fast, slow := head.Next, head
	for fast != end {
		if fast.Val < temp {
			slow = slow.Next
			slow.Val, fast.Val = fast.Val, slow.Val
		}
		fast = fast.Next
	}
	slow.Val, head.Val = head.Val, slow.Val
	quickSort(head, slow)
	quickSort(slow.Next, end)
}
