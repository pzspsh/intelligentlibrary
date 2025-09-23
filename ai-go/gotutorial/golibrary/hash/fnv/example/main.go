package main

import (
	"fmt"
	"hash"
	"hash/fnv"
)

func IncrementalHash(data []byte) uint64 {
	var h hash.Hash64 = fnv.New64a()
	h.Write(data)
	return h.Sum64()
}

func main() {
	data := []byte("hello world")
	hash1 := IncrementalHash(data[:5])
	hash2 := IncrementalHash(data[5:])
	fmt.Println(hash1, hash2)
}
