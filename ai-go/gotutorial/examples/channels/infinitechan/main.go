/*
@File   : main.go
@Author : pan
@Time   : 2024-09-03 16:23:03
*/
package main

// 无限缓存的channel实现
import (
	"fmt"
	"sync"
)

type DynamicChannel struct {
	ch     chan interface{}
	mu     sync.Mutex
	closed bool
}

func NewDynamicChannel(initialCap int) *DynamicChannel {
	return &DynamicChannel{
		ch: make(chan interface{}, initialCap),
	}
}

func (dc *DynamicChannel) Send(value interface{}) {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if dc.closed {
		panic("send on closed channel")
	}

	// 如果channel满了，‌就扩容
	if len(dc.ch) == cap(dc.ch) {
		newCap := cap(dc.ch) * 2
		newCh := make(chan interface{}, newCap)
		for value := range dc.ch {
			newCh <- value
		}
		dc.ch = newCh
	}

	dc.ch <- value
}

func (dc *DynamicChannel) Receive() interface{} {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	if dc.closed && len(dc.ch) == 0 {
		return nil // 或者返回一个错误
	}

	return <-dc.ch
}

func (dc *DynamicChannel) Close() {
	dc.mu.Lock()
	defer dc.mu.Unlock()

	dc.closed = true
	close(dc.ch)
}

func main() {
	dc := NewDynamicChannel(2)
	dc.Send(1)
	dc.Send(2)
	fmt.Println(dc.Receive()) // 输出: 1
	dc.Send(3)
	fmt.Println(dc.Receive()) // 输出: 2
	fmt.Println(dc.Receive()) // 输出: 3
	dc.Close()
	// 下面的接收将返回nil，‌因为channel已经关闭且没有剩余的数据
	fmt.Println(dc.Receive()) // 输出: nil
}
