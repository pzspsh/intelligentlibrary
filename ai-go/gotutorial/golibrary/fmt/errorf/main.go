/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 17:15:50
*/
package main

import (
	"fmt"
)

func main() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}
