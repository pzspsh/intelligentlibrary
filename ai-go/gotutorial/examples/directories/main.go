/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:31:39
*/
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 使用指定的名字和权限创建目录
	err := os.Mkdir("subdir", 0755)
	check(err)
	// 递归删除目录，等效于rm -rf
	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		d := []byte("")
		// 通过写空数据创建文件，文件不存在就创建
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// 强制创建目录，若路径上的父级不存在则创建，等效于mkdir -p
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// 列出指定目录下一级的所有目录和文件
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 改变当前目录，等效于Linux下的cd命令
	err = os.Chdir("subdir/parent/child")
	check(err)
	// 对当前目录执行ReadDir
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	// 递归遍历指定目录下的所有子目录和文件，方式为深度优先遍历
	err = filepath.Walk("subdir", visit)
	if err != nil {
		fmt.Println(err)
	}
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
