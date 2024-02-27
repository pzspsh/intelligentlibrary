/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:32:49
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
