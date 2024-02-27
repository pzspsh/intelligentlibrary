/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:06:01
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //01234
	}
	fmt.Println("")
	fmt.Println(l.Front().Value) //0
	fmt.Println(l.Back().Value)  //4

	l.InsertAfter(6, l.Front()) //首部元素之后插入一个值为6的元素
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //061234
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //601234
	}
	fmt.Println("")
	l.MoveToFront(l.Back()) //将尾部元素移动到首部
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //,460123
	}
	fmt.Println("")
	l2 := list.New()

	for i := 0; i < 5; i++ {
		l2.PushBack(i)
	}
	l2.PushBackList(l) //将l中元素放在l2的末尾
	for e := l2.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //01234460123
	}
	fmt.Println("")
	l.Init()           //清空l
	fmt.Print(l.Len()) //0
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value) //
	}
}
