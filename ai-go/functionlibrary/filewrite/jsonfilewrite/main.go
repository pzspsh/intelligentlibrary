/*
@File   : main.go
@Author : pan
@Time   : 2024-03-01 16:59:42
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 创建要保存的人员信息列表
	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Charlie", 35},
	}

	// 打开或创建 JSON 文件（如果不存在）
	file, err := os.OpenFile("data.json", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 编码并写入 JSON 数据到文件
	encoder := json.NewEncoder(file)
	for _, person := range people {
		err = encoder.Encode(&person)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("成功写入 JSON 文件")
}
