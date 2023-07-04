package main

import "fmt"

func main() {
	first := &ListNode1{Val: 1}
	firsttwo := ListNode1{Val: 2}
	firstthree := ListNode1{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = reverseKGroup1(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func reverseKGroup1(head *ListNode1, k int) *ListNode1 {
	res := 0
	temp := head
	for temp != nil {
		res++
		temp = temp.Next
	}
	if res < k || k <= 1 {
		return head
	}
	pre := &ListNode1{}
	cur := head
	for i := 0; i < k; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	head.Next = reverseKGroup1(cur, k)
	return pre
}
