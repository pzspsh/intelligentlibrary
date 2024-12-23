/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:49:58
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		// sleep 1秒，模拟耗时操作
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// 循环两次，保证从两个channel都读到数据
	for i := 0; i < 2; i++ {
		select {
		// 用select结合case实现多通道等待
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
