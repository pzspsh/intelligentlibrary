/*
@File   : main.go
@Author : pan
@Time   : 2024-05-22 17:09:03
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "path/to/folder" // 替换为你的文件夹路径
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(path) // 输出文件路径
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
