/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:15:50
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")

	io.WriteString(os.Stdout, s) // Ignoring error for simplicity.

}
