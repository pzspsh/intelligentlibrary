/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 12:08:19
*/
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)
}
