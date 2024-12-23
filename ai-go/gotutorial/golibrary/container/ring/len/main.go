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
	// Create a new ring of size 4
	r := ring.New(4)

	// Print out its length
	fmt.Println(r.Len())

}
