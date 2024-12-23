/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 09:57:12
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // Don't forget to flush!
}
