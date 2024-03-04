/*
@File   : main.go
@Author : pan
@Time   : 2024-03-04 15:58:54
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func getRelativePath(file string) {
	currentDir, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		log.Fatal("无法获取当前工作目录")
	}

	absFilePath, err := filepath.Abs(file) // 转换为绝对路径
	if err != nil {
		log.Fatalf("无法获取%s的绝对路径", file)
	}

	relativePath, err := filepath.Rel(currentDir, absFilePath) // 计算相对路径
	if err != nil {
		log.Fatalf("%s不属于当前工作目录或其子目录", file)
	}

	fmt.Println(relativePath)
	f, _ := os.Open(relativePath)
	defer f.Close()
	data, _ := io.ReadAll(f)
	fmt.Println(string(data))
}

func main() {
	getRelativePath("./filepath/file.txt")
}
