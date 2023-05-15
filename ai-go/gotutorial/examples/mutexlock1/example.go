/*
@File   : example.go
@Author : pan
@Time   : 2023-05-12 14:59:25
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	syn := make(chan bool, 10)
	nums := []int{3, 4, 5, 3, 3, 5, 3, 43, 53, 23, 43, 53, 6, 3, 53, 23, 23, 9, 8, 7, 7, 6, 5, 5, 7}
	for index, num := range nums {
		syn <- true
		go func(i, in int) {
			fmt.Printf("第%d 的数据 %d 抢占中\n", in, i)
			mutex.Lock()
			fmt.Printf("第%d 的数据 %d 已占中\n", in, i)
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			<-syn
			fmt.Printf("第%d 的数据 %d 已释放\n", in, i)
		}(num, index)
	}
}
