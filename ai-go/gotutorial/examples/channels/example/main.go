package main

import "fmt"

func main() {
	// 指定缓冲队列最大长度为2
	messages := make(chan string, 2)
	// 发送不再会被阻塞，因为有缓存队列
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
