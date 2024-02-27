/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:20:33
*/
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	r, w, _ := os.Pipe()
	w.Write([]byte("Hello Pipe!"))

	bs := make([]byte, 20)
	n, err := r.Read(bs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bs[:n])) //打印结果： Hello Pipe!
}
