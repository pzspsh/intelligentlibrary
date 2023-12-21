/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:28:02
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	ls, err := os.Lstat("tmp.txt")
	if err != nil {
		log.Fatal(err)
	}

	s, err := os.Stat("tmp.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name:%s,size:%d\n", ls.Name(), ls.Size())
	fmt.Printf("name:%s,size:%d\n", s.Name(), s.Size())
}
