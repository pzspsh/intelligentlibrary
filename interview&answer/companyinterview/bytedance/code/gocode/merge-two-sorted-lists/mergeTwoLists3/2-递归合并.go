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
	node = mergeTwoLists1(&first, &second)

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

// 合并两个有序链表
func mergeTwoLists1(l1 *ListNode1, l2 *ListNode1) *ListNode1 {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}
