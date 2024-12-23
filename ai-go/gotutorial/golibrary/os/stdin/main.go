/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 11:01:17
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type abcReader struct {
}

func (r *abcReader) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

type abcWriter struct {
}

func (r *abcWriter) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func main() {
	var reader abcReader
	var writer abcWriter
	fmt.Printf("please input a string.\n")
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatal("Unable to read/write data")
	}
	/**
	please input a string.
	i love you
	i love you
	*/
}
