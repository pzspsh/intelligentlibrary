/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:32:14
*/
package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int // 定义一个堆，存储结构为数组

// 实现了heap.Interface中组合的sort.Interface接口的Less方法
func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// 实现了heap.Interface中组合的sort.Interface接口的Swap方法
func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// 实现了heap.Interface中组合的sort.Interface接口的Push方法
func (h *myHeap) Len() int {
	return len(*h)
}

// 实现了heap.Interface的Pop方法
func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// 实现了heap.Interface的Push方法
func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

// 按层来遍历和打印堆数据，第一行只有一个元素，即堆顶元素
func (h myHeap) printHeap() {
	n := 1
	levelCount := 1
	for n <= h.Len() {
		fmt.Println(h[n-1 : n-1+levelCount])
		n += levelCount
		levelCount *= 2
	}
}

/*
func main() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap)
	// 用堆本身的Push方法将数组中的元素依次存入堆中
	for _, value := range data {
		aheap.Push(value)
	}

	// 此时堆数组内容为：13, 25, 1, 9, 5, 12, 11
	// 不是正确的堆结构
	aheap.printHeap()
	// 输出：
	//  [13]
	//  [25 1]
	//  [9 5 12 11]

	heap.Init(aheap) // 对堆进行调整，调整后为规范的堆结构

	fmt.Println(*aheap) // 输出：[1 5 11 9 25 12 13]
	aheap.printHeap()
	// 输出：
	//	[1]
	//	[5 11]
	//	[9 25 12 13]
}
*/

/*
func main() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap) // 创建空堆

	// 用heap包中的Push方法将数组中的元素依次存入堆中，
	// 每次Push都会保证堆是规范的堆结构
	for _, value := range data {
		heap.Push(aheap, value)
	}
	fmt.Println(*aheap) // 输出：[1 5 11 25 9 13 12]
	aheap.printHeap()
	// 输出：
	//  [1]
	//  [5 11]
	//  [25 9 13 12]

	// 依次调用heap包的Pop方法来获取堆顶元素
	for aheap.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(aheap))
	}
	// 输出：1 5 9 11 12 13 25
}
*/

/*
func main() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap) // 此时是空堆

	// 用heap包中的Push方法将数组中的元素依次存入堆中，
	// 每次Push都会保证堆是规范的堆结构
	for _, value := range data {
		heap.Push(aheap, value)
		fmt.Printf("插入%d\n", value)
		aheap.printHeap()
		fmt.Println()
	}
	fmt.Println(*aheap)
	// 输出：
	//	插入13
	//	[13]
	//
	//	插入25
	//	[13]
	//	[25]
	//
	//	插入1
	//	[1]
	//	[25 13]
	//
	//	插入9
	//	[1]
	//	[9 13]
	//	[25]
	//
	//	插入5
	//	[1]
	//	[5 13]
	//	[25 9]
	//
	//	插入12
	//	[1]
	//	[5 12]
	//	[25 9 13]
	//
	//	插入11
	//	[1]
	//	[5 11]
	//	[25 9 13 12]
	//
	//	[1 5 11 25 9 13 12]

}
*/

func main() {
	data := [7]int{13, 25, 1, 9, 5, 12, 11}
	aheap := new(myHeap) // 创建空堆

	// 用heap包中的Push方法将数组中的元素依次存入堆中，
	// 每次Push都会保证堆是规范的堆结构
	for _, value := range data {
		heap.Push(aheap, value)
	}
	fmt.Println(*aheap) // 输出：[1 5 11 25 9 13 12]
	aheap.printHeap()
	// 输出：
	//  [1]
	//  [5 11]
	//  [25 9 13 12]

	value := heap.Remove(aheap, 2) // 删除索引号为2的元素（即数组中的第3个元素）
	fmt.Println(value)             // 输出：11
	aheap.printHeap()
	// 输出：
	//	[1]
	//	[5 12]
	//	[25 9 13]
}
