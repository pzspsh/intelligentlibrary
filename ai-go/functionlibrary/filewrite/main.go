/*
@File   : main.go
@Author : pan
@Time   : 2023-06-15 10:00:10
*/
package main

import (
	"log"
	"os"
)

func WriteFile() {
	// 打开文件，‌如果文件不存在则创建，‌存在则以追加的方式打开
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 写入内容
	_, err = file.WriteString("Hello, Go!\n")
	if err != nil {
		log.Fatal(err)
	}

	// 使用fprintf也可以写入内容
	_, err = file.Write([]byte("Another line.\n"))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

}
