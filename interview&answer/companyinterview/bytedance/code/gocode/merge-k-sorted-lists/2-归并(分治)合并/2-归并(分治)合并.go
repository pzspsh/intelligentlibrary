package main

import "fmt"

func main() {
	var first, firstone, firsttwo ListNode1
	first.Val = 1
	firstone.Val = 3
	firsttwo.Val = 5
	first.Next = &firstone
	firstone.Next = &firsttwo

	var second, secondone, secondtwo ListNode1
	second.Val = 2
	secondone.Val = 4
	secondtwo.Val = 6
	second.Next = &secondone
	secondone.Next = &secondtwo

	var node *ListNode1
	node = mergeKLists1([]*ListNode1{&first, &second})

	for {
		fmt.Print(node.Val, "->")
		node = node.Next
		if node.Next == nil {
			fmt.Println(node.Val)
			break
		}
	}
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func mergeKLists1(lists []*ListNode1) *ListNode1 {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	first := mergeKLists1(lists[:len(lists)/2])
	second := mergeKLists1(lists[len(lists)/2:])
	return mergeTwoLists1(first, second)
}

func mergeTwoLists1(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	res := &ListNode1{}
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
