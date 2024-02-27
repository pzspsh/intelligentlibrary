/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:35:33
*/
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string
	Age   int
	Phone []string
}

func main() {
	p := Person{
		Name:  "Bob",
		Age:   18,
		Phone: []string{"123456", "654321"},
	}

	v := reflect.ValueOf(p)
	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", v.Type().Field(i).Name, v.Field(i).Interface())
	}
}
