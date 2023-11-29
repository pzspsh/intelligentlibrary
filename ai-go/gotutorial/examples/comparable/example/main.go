/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 11:11:07
*/
package main

import (
	"fmt"
)

type Dictionay[K comparable, V any] map[K]V

func main() {
	dict := Dictionay[string, int]{"string": 1}
	fmt.Printf("dict: %#v \n", dict)
}
