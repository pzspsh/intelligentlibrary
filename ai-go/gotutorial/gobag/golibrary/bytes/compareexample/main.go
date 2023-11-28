/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:21:39
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Interpret Compare's result by comparing it to zero.
	var a, b []byte
	if bytes.Compare(a, b) < 0 {
		fmt.Println("a less b")
	}
	if bytes.Compare(a, b) <= 0 {
		fmt.Println("a less or equal b")
	}
	if bytes.Compare(a, b) > 0 {
		fmt.Println("a greater b")
	}
	if bytes.Compare(a, b) >= 0 {
		fmt.Println("a greater or equal b")
	}

	// Prefer Equal to Compare for equality comparisons.
	if bytes.Equal(a, b) {
		fmt.Println("a equal b")
	}
	if !bytes.Equal(a, b) {
		fmt.Println("a not equal b")
	}
}
