/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:49:23
*/
package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	logger.Print("Hello, log file!")

	fmt.Print(&buf)
}
