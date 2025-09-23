package main

import (
	"fmt"
)

func main() {
	a := complex(1, 2)
	b := imag(a)
	c := real(a)
	fmt.Println(a, b, c)
}
