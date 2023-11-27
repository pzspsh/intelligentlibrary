/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:31:19
*/
package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}
