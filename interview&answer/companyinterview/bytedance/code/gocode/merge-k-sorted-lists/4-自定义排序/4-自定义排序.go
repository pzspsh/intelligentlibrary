package main

import (
	"fmt"
	"sort"
)

func main() {
	var first, firstone, firsttwo ListNode3
	first.Val = 1
	firstone.Val = 3
	firsttwo.Val = 5
	first.Next = &firstone
	firstone.Next = &firsttwo

	var second, secondone, secondtwo ListNode3
	second.Val = 2
	secondone.Val = 4
	secondtwo.Val = 6
	second.Next = &secondone
	secondone.Next = &secondtwo

	var node *ListNode3
	node = mergeKLists3([]*ListNode3{&first, &second})

	for {
		fmt.Print(node.Val, "->")
		node = node.Next
		if node.Next == nil {
			fmt.Println(node.Val)
			break
		}
	}
}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

func mergeKLists3(lists []*ListNode3) *ListNode3 {
	if len(lists) == 0 {
		return nil
	}
	arr := make([]*ListNode3, 0)
	for i := 0; i < len(lists); i++ {
		temp := lists[i]
		for temp != nil {
			arr = append(arr, temp)
			temp = temp.Next
		}
	}
	if len(arr) == 0 {
		return nil
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})
	for i := 0; i < len(arr)-1; i++ {
		arr[i].Next = arr[i+1]
	}
	arr[len(arr)-1].Next = nil
	return arr[0]
}
