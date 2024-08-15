/*
@File   : main.go
@Author : pan
@Time   : 2024-08-14 14:07:41
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个临时文件作为示例
	tempFile, err := os.CreateTemp("", "example*.txt")
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return
	}
	defer tempFile.Close() // 确保文件在函数结束时关闭

	// 监听SIGINT信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)

	fmt.Println("Program is running. Press Ctrl+C to exit and delete the temp file.")

	// 等待信号
	<-sigChan

	// 接收到SIGINT信号，开始清理
	fmt.Println("Received SIGINT, exiting and deleting the temp file...")

	// 删除文件
	if err := os.Remove(tempFile.Name()); err != nil {
		fmt.Println("Error deleting temp file:", err)
	} else {
		fmt.Println("Temp file deleted successfully.")
	}

	// 退出程序
	fmt.Println("Program exited.")
}
