/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:55:06
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Print(err)
	}
}
