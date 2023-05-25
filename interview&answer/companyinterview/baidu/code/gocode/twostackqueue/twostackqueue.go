/*
@File   : twostackqueue.go
@Author : pan
@Time   : 2023-05-19 14:00:20
*/
package main

func main() {

}

// 用两个栈实现队列
type stack []int

func (s *stack) Push(value int) {
	*s = append(*s, value)
}
func (s *stack) Pop() int {
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

type CQueue struct {
	tail stack
	head stack
}

func Constructor() CQueue {
	return CQueue{}
}

// 1.入队，tail栈保存
// 2.出队, head不为空，出head；head为空，tail出到head里，最后出head
func (c *CQueue) AppendTail(value int) {
	c.tail.Push(value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.head) != 0 {
		return c.head.Pop()
	} else if len(c.tail) != 0 {
		for len(c.tail) > 0 {
			c.head.Push(c.tail.Pop())
		}
		return c.head.Pop()
	}
	return -1
}
