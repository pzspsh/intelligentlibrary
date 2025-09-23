package main

import (
	"bytes"
	"fmt"
)

/*
Repeat返回一个由b的count个副本组成的新字节片。
*/
func main() {
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2)) // banana
	fmt.Println(bytes.Repeat([]byte{1, 2}, 3))        // [1 2 1 2 1 2]
}
