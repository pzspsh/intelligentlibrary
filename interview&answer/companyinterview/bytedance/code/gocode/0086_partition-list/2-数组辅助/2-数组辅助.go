package main

import "fmt"

func main() {
	first := ListNode1{Val: 1}
	firsttwo := ListNode1{Val: 2}
	firstthree := ListNode1{Val: 2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	partition1(&first, 3)

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

func partition1(head *ListNode1, x int) *ListNode1 {
	a := make([]*ListNode1, 0)
	b := make([]*ListNode1, 0)

	for head != nil {
		if head.Val < x {
			a = append(a, head)
		} else {
			b = append(b, head)
		}
		head = head.Next
	}
	temp := &ListNode1{}
	node := temp
	for i := 0; i < len(a); i++ {
		node.Next = a[i]
		node = node.Next
	}
	for i := 0; i < len(b); i++ {
		node.Next = b[i]
		node = node.Next
	}
	node.Next = nil
	return temp.Next
}
