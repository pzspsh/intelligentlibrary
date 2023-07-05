package main

import "fmt"

func main() {
	first := ListNode1{Val: 2}
	first1 := ListNode1{Val: 4}
	//first2 := ListNode{Val: 3}
	first.Next = &first1
	//first1.Next = &first2
	fmt.Println(rotateRight1(&first, 1))
}

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func rotateRight1(head *ListNode1, k int) *ListNode1 {
	if head == nil || k == 0 {
		return head
	}
	temp := head
	count := 0
	arr := make([]*ListNode1, 0)
	for temp != nil {
		arr = append(arr, temp)
		temp = temp.Next
		count++
	}
	k = k % count
	if k == 0 {
		return head
	}
	arr[count-1].Next = head
	temp = arr[count-1-k]
	head, temp.Next = temp.Next, nil
	return head
}
