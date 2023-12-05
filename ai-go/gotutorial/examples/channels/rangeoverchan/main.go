/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:53:05
*/
package main

import "fmt"

/* 使用range也可用于从通道接收消息，包括在已关闭的通道上。 */
func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// 即使通道已关闭，仍然可以用range接收消息
	for elem := range queue {
		fmt.Println(elem)
	}
}
