/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:15:50
*/
package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// It is conventional not to worry about any
	// error returned by Printf.
}
