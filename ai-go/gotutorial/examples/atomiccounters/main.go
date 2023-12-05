/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:59:30
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/* Go中的同步方式除了通过channel通信，还可以通过原子计数器。通过原子技术器访问或者修改数值变量，可以让这些操作成为原子性的。 */
func main() {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// 通过AddUnit64实现原子性增加，若是原子读取可用LoadUint64，注意这里需要传指针参数
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}
