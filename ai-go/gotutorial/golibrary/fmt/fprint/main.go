/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:15:50
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// The n and err return values from Fprint are
	// those returned by the underlying io.Writer.
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

}
