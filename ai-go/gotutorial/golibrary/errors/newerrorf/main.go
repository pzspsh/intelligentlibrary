/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 16:55:06
*/
package main

import (
	"fmt"
)

func main() {
	const name, id = "bimmler", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	if err != nil {
		fmt.Print(err)
	}
}
