package main

import "fmt"

func main() {
	var first, firstone, firsttwo ListNode2
	first.Val = 1
	firstone.Val = 3
	firsttwo.Val = 5
	first.Next = &firstone
	firstone.Next = &firsttwo

	var second, secondone, secondtwo ListNode2
	second.Val = 2
	secondone.Val = 4
	secondtwo.Val = 6
	second.Next = &secondone
	secondone.Next = &secondtwo

	var node *ListNode2
	node = mergeTwoLists2(&first, &second)

	for {
		fmt.Print(node.Val, "->")
		node = node.Next
		if node.Next == nil {
			fmt.Println(node.Val)
			break
		}
	}
}

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func mergeTwoLists2(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	res := &ListNode2{}
	temp := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}
	if l1 != nil {
		temp.Next = l1
	} else {
		temp.Next = l2
	}
	return res.Next
}
