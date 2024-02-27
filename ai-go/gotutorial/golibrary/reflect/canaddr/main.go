/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:34:05
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "hello reflection"
	value := reflect.ValueOf(&str)

	fmt.Println("CanAddr:", value.CanAddr())
}
