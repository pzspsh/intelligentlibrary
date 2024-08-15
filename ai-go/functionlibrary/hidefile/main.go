/*
@File   : main.go
@Author : pan
@Time   : 2024-08-14 14:39:39
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	// 定义隐藏文件的名称，以'.'开头
	hiddenFileName := ".hiddenfile.txt"

	// 打开（如果不存在则创建）一个文件用于写入
	// 注意：这里的O_CREATE和O_WRONLY标志在os.OpenFile中不是直接使用的，
	// 而是通过os.O_CREATE | os.O_WRONLY这样的组合来使用。
	// 但为了简单起见，这里使用os.OpenFile的便捷函数os.Create。
	hiddenFile, err := os.Create(hiddenFileName)
	if err != nil {
		fmt.Println("Error creating hidden file:", err)
		return
	}
	defer hiddenFile.Close() // 确保在函数结束时关闭文件

	// 向文件中写入一些内容（可选）
	_, err = hiddenFile.WriteString("This is a hidden file.\n")
	if err != nil {
		fmt.Println("Error writing to hidden file:", err)
		return
	}

	fmt.Println("Hidden file created successfully.")
}
