package main

import "fmt"

func main() {
	first := &ListNode3{Val: 1}
	firsttwo := ListNode3{Val: 2}
	firstthree := ListNode3{Val: 3}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	firstthree.Next = &ListNode3{Val: 4}
	first = reverseKGroup3(first, 3)
	for first != nil {
		fmt.Println(first.Val)
		first = first.Next
	}
}

type ListNode3 struct {
	Val  int
	Next *ListNode3
}

// leetcode25_K个一组翻转链表
func reverseKGroup3(head *ListNode3, k int) *ListNode3 {
	res := &ListNode3{Next: head}
	prev, end := res, res
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := prev.Next         // 开始的位置
		next := end.Next           // 结束的下一个位置
		end.Next = nil             // 断开尾部连接
		prev.Next = reverse3(start) // 反转后接到prev.Next
		start.Next = next          // start的指针指向下一个开头（此时start已经是反转的最后一个节点）
		prev = start               // 已经处理后的最后一个节点
		end = prev                 // end也移动到prev
	}
	return res.Next
}

func reverse3(head *ListNode3) *ListNode3 {
	var result *ListNode3
	for head != nil {
		temp := head.Next
		head.Next = result
		result = head
		head = temp
	}
	return result
}
