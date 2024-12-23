package main

import "fmt"

func main() {
	first := &ListNode1{Val: 1}
	firsttwo := ListNode1{Val: 2}
	firstthree := ListNode1{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = swapPairs1(first)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func swapPairs1(head *ListNode1) *ListNode1 {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	first.Next, second.Next = swapPairs1(second.Next), first
	return second
}
