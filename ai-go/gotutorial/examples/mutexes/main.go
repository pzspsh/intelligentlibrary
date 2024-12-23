/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:00:44
*/
package main

import (
	"fmt"
	"sync"
)

// 封装结构，包括：互斥锁和它保护的map计数器
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// 通过互斥锁保证对计数器的修改是原子操作
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

/* Go中处理同步的另一方式是使用互斥锁，这也是许多其他语言会采用的方式，保证加锁的代码块为原子操作。 */
func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
