/*
@File   : main.go
@Author : pan
@Time   : 2024-09-03 16:24:22
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个初始缓冲为2的channel
	messages := make(chan string, 2)

	// 这个goroutine会将消息发送到channel
	go func() {
		for i := 0; ; i++ {
			message := fmt.Sprintf("message %d", i)
			messages <- message
			time.Sleep(1 * time.Second)
		}
	}()

	// 这个goroutine会从channel接收消息
	go func() {
		for {
			message := <-messages
			fmt.Println(message)
			// 当缓冲满了之后，我们增加缓冲的大小
			if len(messages) == cap(messages) {
				newCap := cap(messages) * 2
				newMessages := make(chan string, newCap)
				go func() {
					for msg := range messages {
						newMessages <- msg
					}
				}()
				messages = newMessages
			}
		}
	}()
	// 阻塞主goroutine
	select {}
}
