/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:28:15
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var in float32 = 3.1415926
	out := (reflect.ValueOf(complex64(complex(in, 0)))).Interface().(complex64)
	fmt.Printf("Input: %v, Output: %v, Type: %v\n", in, out, reflect.TypeOf(out))

	var in2 float64 = 3.1415926535
	out2 := (reflect.ValueOf(complex(in2, 0))).Interface().(complex128)
	fmt.Printf("Input: %v, Output: %v, Type: %v\n", in2, out2, reflect.TypeOf(out2))
}
