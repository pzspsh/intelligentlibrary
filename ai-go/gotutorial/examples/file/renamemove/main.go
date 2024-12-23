/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:30:35
*/
package main

import (
	"log"
	"os"
)

func main() {
	originalPath := "test.txt"
	newPath := "test2.txt"
	err := os.Rename(originalPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
