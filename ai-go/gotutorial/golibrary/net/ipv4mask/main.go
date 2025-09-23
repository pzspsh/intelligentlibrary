package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.IPv4Mask(255, 255, 255, 0))
}
