/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:49:02
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a =", a) // a = [0 1 2 6]

	// ...

	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))
	*pname = "Alice"
	*page = 28
	*pgender = false
	fmt.Println(who) // {Alice 28 false}
}
