package main

import (
	"fmt"
	"net"
)

func main() {
	hw, error := net.ParseMAC("0123.4567.89ab.cdef.0000.0123.4567.89ab.cdef.0000")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(hw.String())
}
