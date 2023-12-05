/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:48:15
*/
package main

import "fmt"

func main() {
	// channel是引用类型，需要通过make创建
	messages := make(chan string)

	// 发送ping消息，并阻塞等待接收方
	go func() { messages <- "ping" }()
	// 接收ping消息，在接收到消息前阻塞
	msg := <-messages
	fmt.Println(msg)
}
