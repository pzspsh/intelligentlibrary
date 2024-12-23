/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:42:12
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/* 下面例子展示了信号量的使用。通过监听中止信号量，在外部强行中断程序后，还能执行一段结束代码。 */
func main() {
	sigs := make(chan os.Signal, 1)
	// 监听程序中止的信号量，并存入通道sigs
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	go func() {
		// 从sigs通道读取信号量并打印
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
