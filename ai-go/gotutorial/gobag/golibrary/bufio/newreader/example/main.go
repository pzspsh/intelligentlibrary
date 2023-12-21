/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:14:12
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D:/GoLang/go_demo/helloworld/hello.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //必须关闭文件流
	//bufio读取文件
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //表示一次读取一行
		if err == io.EOF {
			fileStr += line
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += line
	}
	fmt.Println(fileStr)
}
