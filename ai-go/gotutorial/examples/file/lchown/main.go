/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:35:38
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 创建一个硬链接。
	// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
	// 删除和重命名不会影响另一个。
	err := os.Link("original.txt", "original_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("creating sym")
	// Create a symlink
	err = os.Symlink("original.txt", "original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
	// Symlink在Windows中不工作。
	fileInfo, err := os.Lstat("original_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	//改变软链接的拥有者不会影响原始文件。
	err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
}
