/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 16:12:20
*/
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	// Iterate through the ring and print its contents
	for j := 0; j < n; j++ {
		fmt.Println(r.Value)
		r = r.Next()
	}

}
