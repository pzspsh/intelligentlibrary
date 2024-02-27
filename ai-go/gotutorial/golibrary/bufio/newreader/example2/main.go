/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 17:08:26
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, _ := os.Open("file.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if r == '\n' {
			fmt.Println()
		} else {
			fmt.Print(r)
		}
	}
}
