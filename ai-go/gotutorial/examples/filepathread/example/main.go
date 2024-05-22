/*
@File   : main.go
@Author : pan
@Time   : 2024-05-22 17:10:57
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	root := "path/to/folder" // 替换为你的文件夹路径
	readDir(root, "")
}

func readDir(path string, indent string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	f.Close()

	for _, file := range files {
		if file.IsDir() {
			newPath := path + "/" + file.Name()
			fmt.Println(indent + newPath) // 输出文件夹路径
			readDir(newPath, indent+"  ")
		} else {
			fmt.Println(indent + file.Name()) // 输出文件路径
		}
	}
}
