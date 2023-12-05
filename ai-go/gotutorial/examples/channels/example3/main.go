/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:49:04
*/
package main

import "fmt"

// 先发送消息给pings
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 再从pings中取消息，发送给pongs
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	// 最后从pongs中取消息
	fmt.Println(<-pongs)
}
