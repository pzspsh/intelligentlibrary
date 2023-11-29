/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:31:03
*/
package main

import (
	"container/list"
	"fmt"
)

// 打印链表
func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Printf("\n")
}

func main() {
	list := list.New()
	two := list.PushBack(2)
	three := list.PushBack(3)
	one := list.PushFront(1)
	printList(list) // 1 2 3

	list.MoveToBack(one)    // 2 3 1
	list.MoveToFront(three) // 3 2 1
	printList(list)         // 3 2 1
	list.Remove(two)
	printList(list) //3 1

	list.InsertAfter(4, one)
	list.InsertBefore(5, three)
	printList(list) // 5 3 1 4

	println(list.Len()) // 4
	list.Init()
	println(list.Len()) // 0
}
