package main

import "fmt"

func main() {
	first := &ListNode2{Val: 1}
	firsttwo := ListNode2{Val: 2}
	firstthree := ListNode2{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = removeNthFromEnd2(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

var count int

func removeNthFromEnd2(head *ListNode2, n int) *ListNode2 {
	if head == nil {
		count = 0
		return nil
	}
	head.Next = removeNthFromEnd2(head.Next, n)
	count = count + 1
	if count == n {
		return head.Next
	}
	return head
}
