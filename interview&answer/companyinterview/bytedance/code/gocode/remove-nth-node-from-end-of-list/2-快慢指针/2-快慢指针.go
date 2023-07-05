package main

import "fmt"

func main() {
	first := &ListNode1{Val: 1}
	firsttwo := ListNode1{Val: 2}
	firstthree := ListNode1{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	first = removeNthFromEnd1(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// leetcode19_删除链表的倒数第N个节点
func removeNthFromEnd1(head *ListNode1, n int) *ListNode1 {
	temp := &ListNode1{Next: head}
	fast, slow := temp, temp
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return temp.Next
}
