package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var n int64 = 5
	var pn = &n
	var pf = (*float64)(unsafe.Pointer(pn))
	// now, pn and pf are pointing at the same memory address
	fmt.Println(*pf) // 2.5e-323
	*pf = 3.14159
	fmt.Println(n) // 4614256650576692846
}
