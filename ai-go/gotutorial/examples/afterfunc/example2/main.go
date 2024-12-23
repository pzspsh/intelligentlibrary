/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:19:18
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始执行程序")
	duration := time.Second * 1
	timer := time.AfterFunc(duration, func() {
		fmt.Println("1秒钟后执行该函数")
	})
	for i := 0; i < 10; i++ {
		<-time.After(duration)
		fmt.Println("3秒钟后执行该函数")
		timer.Reset(duration)
	}
	fmt.Println("程序结束")
}
