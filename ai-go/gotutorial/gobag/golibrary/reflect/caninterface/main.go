/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:32:09
*/
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Jack", Age: 30}
	value := reflect.ValueOf(p)
	fmt.Println(value.CanInterface()) // true

	a := 1
	value1 := reflect.ValueOf(a)
	fmt.Println(value1.CanInterface()) // true
}
