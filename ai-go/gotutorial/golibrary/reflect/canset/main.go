package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.1415
	v := reflect.ValueOf(x)
	fmt.Println("v.CanSet() =", v.CanSet())

	y := reflect.ValueOf(&x).Elem()
	fmt.Println("y.CanSet() =", y.CanSet())

	z := reflect.ValueOf(x)
	fmt.Println("z.CanAddr() =", z.CanAddr())
}
