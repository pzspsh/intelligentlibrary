package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
}
