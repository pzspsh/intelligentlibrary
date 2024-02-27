/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:14:41
*/
package main

import (
	"fmt"
	"index/suffixarray"
)

func main() {
	index := suffixarray.New([]byte("banana"))
	offsets := index.Lookup([]byte("ana"), -1)
	for _, off := range offsets {
		fmt.Println(off)
	}

}
