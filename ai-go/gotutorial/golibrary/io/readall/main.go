/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:39:38
*/
package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
