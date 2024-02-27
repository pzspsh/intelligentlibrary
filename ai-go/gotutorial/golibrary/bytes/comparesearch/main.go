/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:23:15
*/
package main

import (
	"bytes"
	"fmt"
	"sort"
)

func main() {
	// Binary search to find a matching byte slice.
	var needle []byte
	var haystack [][]byte // Assume sorted
	i := sort.Search(len(haystack), func(i int) bool {
		// Return haystack[i] >= needle.
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		fmt.Println("Found it!")
	}
}
