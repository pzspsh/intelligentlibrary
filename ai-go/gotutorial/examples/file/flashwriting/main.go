/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:40:56
*/
package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
