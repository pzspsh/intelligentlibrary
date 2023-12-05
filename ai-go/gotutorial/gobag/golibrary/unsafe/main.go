/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:49:24
*/
package main

import (
	"fmt"
	"unsafe"
)

// case A: conversions between unsafe.Pointer and uintptr
//
//	don't appear in the same expression
func illegalUseA() {
	fmt.Println("===================== illegalUseA")

	pa := new([4]int)

	// split the legal use
	// p1 := unsafe.Pointer(uintptr(unsafe.Pointer(pa)) + unsafe.Sizeof(pa[0]))
	// into two expressions (illegal use):
	ptr := uintptr(unsafe.Pointer(pa))
	p1 := unsafe.Pointer(ptr + unsafe.Sizeof(pa[0]))
	// "go vet" will make a warning for the above line:
	// possible misuse of unsafe.Pointer

	// the unsafe package docs, https://golang.org/pkg/unsafe/#Pointer,
	// thinks above splitting is illegal.
	// but the current Go compiler and runtime (1.7.3) can't detect
	// this illegal use.
	// however, to make your program run well for later Go versions,
	// it is best to comply with the unsafe package docs.

	*(*int)(p1) = 123
	fmt.Println("*(*int)(p1)  :", *(*int)(p1)) //
}

// case B: pointers are pointing at unknown addresses
func illegalUseB() {
	fmt.Println("===================== illegalUseB")

	a := [4]int{0, 1, 2, 3}
	p := unsafe.Pointer(&a)
	p = unsafe.Pointer(uintptr(p) + uintptr(len(a))*unsafe.Sizeof(a[0]))
	// now p is pointing at the end of the memory occupied by value a.
	// up to now, although p is invalid, it is no problem.
	// but it is illegal if we modify the value pointed by p
	*(*int)(p) = 123
	fmt.Println("*(*int)(p)  :", *(*int)(p)) // 123 or not 123
	// the current Go compiler/runtime (1.7.3) and "go vet"
	// will not detect the illegal use here.

	// however, the current Go runtime (1.7.3) will
	// detect the illegal use and panic for the below code.
	p = unsafe.Pointer(&a)
	for i := 0; i <= len(a); i++ {
		*(*int)(p) = 123 // Go runtime (1.7.3) never panic here in the tests

		fmt.Println(i, ":", *(*int)(p))
		// panic at the above line for the last iteration, when i==4.
		// runtime error: invalid memory address or nil pointer dereference

		p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(a[0]))
	}
}

func main() {
	illegalUseA()
	illegalUseB()
}
