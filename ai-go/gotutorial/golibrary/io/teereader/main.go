/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:39:38
*/
package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var r io.Reader = strings.NewReader("some io.Reader stream to be read\n")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}

}
