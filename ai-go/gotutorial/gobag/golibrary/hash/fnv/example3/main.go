/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 10:14:06
*/
package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	const MASK uint64 = 1<<63 - 1
	h := fnv.New64()
	h.Write([]byte("1133"))
	hash := h.Sum64()
	fmt.Printf("%#x\n", MASK)
	fmt.Println(hash)
	hash = (hash >> 63) ^ (hash & MASK)
	fmt.Println(hash)
}
