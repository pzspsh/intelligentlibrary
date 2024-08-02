/*
@File   : main.go
@Author : pan
@Time   : 2024-08-02 11:01:44
*/
package main

import (
	"fmt"
)

func main() {
	for i := range 100 {
		fmt.Println("hello: ", i)
	}
	fmt.Println("go 1.22 has lift-off!")
}
