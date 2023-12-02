/*
@File   : main.go
@Author : pan
@Time   : 2023-12-02 20:39:19
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// This example shows a case in which the name of a promoted field
	// is hidden by another field: FieldByName will not work, so
	// FieldByIndex must be used instead.
	type user struct {
		firstName string
		lastName  string
	}

	type data struct {
		user
		firstName string
		lastName  string
	}

	u := data{
		user:      user{"Embedded John", "Embedded Doe"},
		firstName: "John",
		lastName:  "Doe",
	}

	s := reflect.ValueOf(u).FieldByIndex([]int{0, 1})
	fmt.Println("embedded last name:", s)
}
