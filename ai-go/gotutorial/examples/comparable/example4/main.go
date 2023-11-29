/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 11:50:30
*/
package main

import (
	"fmt"
)

type Addable interface {
	int
	interface{}
}

func add[T Addable](a T) T {
	return a
}

func main() {
	fmt.Println(add(1))
}
