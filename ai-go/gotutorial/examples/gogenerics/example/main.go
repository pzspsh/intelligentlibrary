package main

import (
	"fmt"
)

type Dictionay[K comparable, V any] map[K]V

func main() {
	dict := Dictionay[string, int]{"string": 1}
	fmt.Printf("dict: %#v \n", dict)
}
