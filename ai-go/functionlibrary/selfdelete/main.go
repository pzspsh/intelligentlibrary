/*
@File   : main.go
@Author : pan
@Time   : 2024-09-11 16:31:24
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func SelfDelete() error {
	var err error
	path, err := os.Executable()
	if err != nil {
		return err
	}
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "del", path)
		if err = cmd.Start(); err != nil { // 修复了错误处理
			return err
		}
	} else {
		if err = os.Remove(path); err != nil {
			if err = os.Chmod(path, 0o775); err != nil {
				return err
			} else {
				if err = os.Remove(path); err != nil {
					return err
				}
			}
		}
	}
	return err
}

func main() {
	// 程序的主要逻辑
	fmt.Println("程序的主要逻辑")
	// 程序退出前尝试删除自身
	if err := SelfDelete(); err != nil {
		// 日志错误或其他处理方式
		fmt.Println("无法删除自身:", err.Error())
	}
}
