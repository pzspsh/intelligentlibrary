/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:36:45
*/
package main

import (
	"compress/bzip2"
	"fmt"
	"log"
	"os"
)

func main() {
	fz, err := os.Open("filepath/test.bz2")
	if err != nil {
		log.Fatal(err)
	}
	w := bzip2.NewReader(fz)
	buf := make([]byte, 1024*100)
	for {
		n, err := w.Read(buf)
		if n == 0 || err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
