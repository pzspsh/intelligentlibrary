/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:47:40
*/
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	// 可以通过go执行函数来启动协程
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 注意协程不会阻止主线程的结束，即使协程还未执行完。所以这里需要sleep 1秒
	time.Sleep(time.Second)
	fmt.Println("done")
}
