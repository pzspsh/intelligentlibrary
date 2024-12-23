/*
@File   : folderexists.go
@Author : pan
@Time   : 2023-06-21 11:06:59
*/
package main

import (
	"fmt"
	"os"
)

func FolderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		return false
	}
	return info.IsDir()
}

func main() {
	result := FolderExists("../tests")
	fmt.Println(result)
}
