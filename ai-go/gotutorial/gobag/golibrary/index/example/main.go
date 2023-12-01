/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:14:41
*/
package main

import (
	"fmt"
	"index/suffixarray"
	"sort"
)

func main() {
	source := []byte("hello world, hello china")
	index := suffixarray.New(source)
	offsets := index.Lookup([]byte("hello"), -1)
	sort.Ints(offsets)
	fmt.Printf("%v", offsets)
}
