package main

import "fmt"

func main() {
	first := ListNode1{Val: 1}
	firsttwo := ListNode1{Val: 2}
	firstthree := ListNode1{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	deleteDuplicates1(&first)

	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = *first.Next
	}
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func deleteDuplicates1(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicates1(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}
