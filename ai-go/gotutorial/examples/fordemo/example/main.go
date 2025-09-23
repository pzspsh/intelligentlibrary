package main

import (
	"fmt"
)

func main() {
	for i := range 100 {
		fmt.Println("hello: ", i)
	}
	fmt.Println("go 1.22 has lift-off!")
}
