package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type MyInt int

	a := []MyInt{0, 1, 2}
	// b := ([]int)(a) // error: cannot convert a (type []MyInt) to type []int
	b := *(*[]int)(unsafe.Pointer(&a))

	b[0] = 3

	fmt.Println("a =", a) // a = [3 1 2]
	fmt.Println("b =", b) // b = [3 1 2]

	a[2] = 9

	fmt.Println("a =", a) // a = [3 1 9]
	fmt.Println("b =", b) // b = [3 1 9]
}
