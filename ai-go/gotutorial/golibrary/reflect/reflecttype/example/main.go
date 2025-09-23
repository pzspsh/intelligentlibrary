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
	p := Person{"John", 30}
	t := reflect.TypeOf(p)

	fmt.Println(t.Name())        // Person
	fmt.Println(t.Kind())        // struct
	fmt.Println(t.NumField())    // 2
	fmt.Println(t.Field(0).Name) // Name
	fmt.Println(t.Field(0).Type) // string
	fmt.Println(t.Field(1).Name) // Age
	fmt.Println(t.Field(1).Type) // int
}
