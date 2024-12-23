/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 22:03:25
*/
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Alice", "Bob", "Vera"}
	fmt.Println("Equal:", slices.Compare(names, []string{"Alice", "Bob", "Vera"}))
	fmt.Println("V < X:", slices.Compare(names, []string{"Alice", "Bob", "Xena"}))
	fmt.Println("V > C:", slices.Compare(names, []string{"Alice", "Bob", "Cat"}))
	fmt.Println("3 > 2:", slices.Compare(names, []string{"Alice", "Bob"}))
}
