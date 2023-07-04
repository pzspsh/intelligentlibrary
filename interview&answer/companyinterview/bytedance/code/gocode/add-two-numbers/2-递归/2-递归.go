package main

import "fmt"

func main() {
	first := ListNode1{Val: 2}
	first1 := ListNode1{Val: 4}
	first2 := ListNode1{Val: 3}
	first.Next = &first1
	first1.Next = &first2

	second := ListNode1{Val: 5}
	second1 := ListNode1{Val: 6}
	second2 := ListNode1{Val: 4}
	second.Next = &second1
	second1.Next = &second2

	fmt.Println(addTwoNumbers1(&first, &second))
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func addTwoNumbers1(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum := l1.Val + l2.Val
	res := &ListNode1{Val: sum % 10}
	if sum >= 10 {
		l1.Next = addTwoNumbers1(l1.Next, &ListNode1{Val: 1})
	}
	res.Next = addTwoNumbers1(l1.Next, l2.Next)
	return res
}
