/*
@File   : createfolders.go
@Author : pan
@Time   : 2023-06-21 15:00:53
*/
package main

import (
	"fmt"
	"os"
)

func CreateFolders(paths ...string) error {
	for _, path := range paths {
		if err := CreateFolder(path); err != nil {
			return err
		}
	}

	return nil
}

func CreateFolder(path string) error {
	return os.MkdirAll(path, 0700)
}

func main() {
	err := CreateFolders("../tests/test1", "../tests/test2")
	if err != nil {
		fmt.Println("err:", err)
	}
}
