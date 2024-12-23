/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:45:32
*/
package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	root := "/usr/local/go/bin"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}
