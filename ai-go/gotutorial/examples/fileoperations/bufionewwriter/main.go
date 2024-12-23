/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:32:05
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello word\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}
