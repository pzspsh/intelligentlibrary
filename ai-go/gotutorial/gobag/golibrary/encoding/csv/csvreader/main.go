/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:15:19
*/
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("filepath/data.csv")
	if err != nil {
		fmt.Println("文件打开失败: ", err)
		return
	}
	// 延迟关闭文件
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}

		if err != nil {
			fmt.Println("读取文件时发生错误: ", err)
			return
		}
		fmt.Println("该行内容为: ", line)
	}
}
