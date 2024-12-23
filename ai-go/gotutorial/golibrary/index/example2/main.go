/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 12:16:46
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"index/suffixarray"
	"strings"
)

func main() {
	index := suffixarray.New([]byte("banana"))

	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	index.Write(writer)
	writer.Flush()

	rIndex := suffixarray.New(nil)
	rIndex.Read(bytes.NewReader(b.Bytes()))
	offsets := rIndex.Lookup([]byte("ana"), -1)
	for idx, off := range offsets {
		fmt.Printf("suffixarryay%d: %d\n", idx, off)
	}

	fmt.Println("strings.Index1: ", strings.Index("banana", "ana"))
	fmt.Println("strings.Index2: ", strings.Index("banana", "ana"))
}
