package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始执行程序")
	timer := time.AfterFunc(time.Second*5, func() {
		fmt.Println("5秒后执行该函数")
	})
	time.Sleep(time.Second * 2)
	timer.Stop()
	fmt.Println("已经停止执行函数")
	time.Sleep(time.Second * 10)
	fmt.Println("程序结束")
}
