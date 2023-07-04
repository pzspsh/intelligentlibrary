package main

import "fmt"

func main() {
	first := &ListNode2{Val: 1}
	firsttwo := ListNode2{Val: 2}
	firstthree := ListNode2{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = reverseKGroup2(first, 2)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func reverseKGroup2(head *ListNode2, k int) *ListNode2 {
	res := &ListNode2{Next: head}
	prev := res
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return res.Next
			}
		}
		next := tail.Next
		head, tail = reverse2(head, tail)
		prev.Next = head
		tail.Next = next
		prev = tail
		head = tail.Next
	}
	return res.Next
}

func reverse2(head, tail *ListNode2) (*ListNode2, *ListNode2) {
	prev := tail.Next
	temp := head
	for prev != tail {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return tail, head
}
