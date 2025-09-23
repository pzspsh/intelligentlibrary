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
