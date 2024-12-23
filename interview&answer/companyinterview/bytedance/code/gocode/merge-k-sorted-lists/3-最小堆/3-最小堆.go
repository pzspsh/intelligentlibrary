package main

import (
	"container/heap"
	"fmt"
)

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
	node = mergeKLists2([]*ListNode2{&first, &second})

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

func mergeKLists2(lists []*ListNode2) *ListNode2 {
	if len(lists) == 0 {
		return nil
	}
	var h IntHeap
	heap.Init(&h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(&h, lists[i])
		}
	}
	res := &ListNode2{}
	temp := res
	for h.Len() > 0 {
		minItem := heap.Pop(&h).(*ListNode2)
		temp.Next = minItem
		temp = temp.Next
		if minItem.Next != nil {
			heap.Push(&h, minItem.Next)
		}
	}
	return res.Next
}

type IntHeap []*ListNode2

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode2)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
