package main

import "fmt"

func main() {
	first := ListNode2{Val: 1}
	firsttwo := ListNode2{Val: 2}
	firstthree := ListNode2{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	deleteDuplicates2(&first)

	for {
		fmt.Println(first.Val)
		if first.Next == nil {
			break
		}
		first = *first.Next
	}
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func deleteDuplicates2(head *ListNode2) *ListNode2 {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := head.Next
	for p.Next != nil {
		if p.Val == q.Val {
			if q.Next == nil {
				p.Next = nil
			} else {
				p.Next = q.Next
				q = q.Next
			}
		} else {
			p = p.Next
			q = q.Next
		}
	}
	return head
}
