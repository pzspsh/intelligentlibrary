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
