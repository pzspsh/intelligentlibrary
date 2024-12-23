/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:38:32
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("test.txt")
	defer file.Close()

	// 偏离位置，可以是正数也可以是负数
	var offset int64 = 5

	// 用来计算offset的初始位置
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	var whence int = 0
	newPosition, err := file.Seek(offset, whence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved to 5:", newPosition)

	// 从当前位置回退两个字节
	newPosition, err = file.Seek(-2, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Just moved back two:", newPosition)

	// 使用下面的技巧得到当前的位置
	currentPosition, err := file.Seek(0, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current position:", currentPosition)

	// 转到文件开始处
	newPosition, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Position after seeking 0,0:", newPosition)
}
