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
	fmt.Println(name, "is", age, "years old.")

	// It is conventional not to worry about any
	// error returned by Println.
}
