/*
@File   : main.go
@Author : pan
@Time   : 2024-03-01 17:06:14
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 创建要添加的新人员信息
	newPerson := &Person{
		Name: "John",
		Age:  31,
	}

	// 将新人员信息转换为JSON格式字符串
	data, err := json.Marshal(newPerson)
	if err != nil {
		log.Fatal("Failed to marshal JSON data: ", err)
	}

	// 打开现有的JSON文件（如果不存在则会自动创建）
	filePath := "file.json" // 更改为你的文件路径
	// os.O_APPEND 写入时向文件追加数据
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Failed to open file for writing: ", err)
	}
	defer file.Close()

	/*
	   // 向文件中追加写入JSON数据
	   	_, err = file.WriteString("\n") // 先写入换行符分隔每条记录
	   	if err != nil {
	   		log.Fatal("Failed to write newline character: ", err)
	   	}
	*/
	_, err = file.Write(append(data, []byte{'\n'}...)) // 再写入JSON数据前需要添加制表符作为缩进
	if err != nil {
		log.Fatal("Failed to append JSON data: ", err)
	}

	fmt.Println("成功追加写入JSON文件！")
}
