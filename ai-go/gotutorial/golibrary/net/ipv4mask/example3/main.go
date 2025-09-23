package main

import (
	"fmt"
	"net"
)

func main() {
	mask := net.IPMask(net.ParseIP("255.255.255.0").To4()) // If you have the mask as a string
	//mask := net.IPv4Mask(255,255,255,0) // If you have the mask as 4 integer values

	prefixSize, _ := mask.Size()
	fmt.Println(prefixSize)
}
