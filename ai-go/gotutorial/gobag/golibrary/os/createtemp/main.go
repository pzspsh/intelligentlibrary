/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 18:45:19
*/
package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
