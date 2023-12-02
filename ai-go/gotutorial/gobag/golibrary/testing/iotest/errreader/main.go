/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 23:45:13
*/
package main

import (
	"errors"
	"fmt"
	"testing/iotest"
)

func main() {
	// A reader that always returns a custom error.
	r := iotest.ErrReader(errors.New("custom error"))
	n, err := r.Read(nil)
	fmt.Printf("n:   %d\nerr: %q\n", n, err)

}
