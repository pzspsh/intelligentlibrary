/*
@File   : main.go
@Author : pan
@Time   : 2024-09-11 16:31:24
*/
package main

import (
	"os"
)

func SelfDelete() error {
	path, err := os.Executable()
	if err != nil {
		return err
	}
	return os.Remove(path)
}

func main() {
	// 程序的主要逻辑

	// 程序退出前尝试删除自身
	if err := SelfDelete(); err != nil {
		// 日志错误或其他处理方式
		println("无法删除自身:", err.Error())
	}
}
