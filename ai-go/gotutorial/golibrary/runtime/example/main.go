/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 21:46:33
*/
package main

// 判断操作系统

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("操作系统是 macOS。")
	case "linux":
		fmt.Println("操作系统是 Linux。")
	case "windows":
		fmt.Println("操作系统是 Windows。")
	default:
		// 如果不是上面列出的系统，则打印未知
		fmt.Printf("未知的操作系统: %s\n", os)
	}
}
