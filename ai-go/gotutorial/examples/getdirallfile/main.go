/*
@File   : main.go
@Author : pan
@Time   : 2024-08-20 11:17:02
*/
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("error:", err)
			return err
		}
		*files = append(*files, path)
		return nil
	}
}

func GetAllFile(dir string) {
	var files []string
	// dir 设置为目标目录
	err := filepath.Walk(dir, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

func GetAllFile1(dir string) {
	// dir 设置为目标目录
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
		fmt.Println(filepath.Join(dir, entry.Name()))
	}
}

func GetAllFile2(dir string) {
	var err error
	if err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() { // 处理文件
			filename := filepath.Base(path)
			fmt.Println("filename:", filename, "path:", path)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// GetAllFile("path/intelligentlibrary/docs")
	// GetAllFile1("path/intelligentlibrary/docs")
	// GetAllFile2("path/intelligentlibrary")
}
