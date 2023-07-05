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
	temp := &ListNode2{Next: head}
	second := temp
	first := head
	for first != nil && first.Next != nil {
		if first.Next.Val != second.Next.Val {
			first = first.Next
			second = second.Next
		} else {
			for first.Next != nil && first.Next.Val == second.Next.Val {
				first = first.Next
			}
			second.Next = first.Next
			first = first.Next
		}
	}
	return temp.Next
}
