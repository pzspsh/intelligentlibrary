/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:27:22
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	str := "Hello World"
	dst := reflect.New(reflect.ValueOf(str).Type()).Elem()
	src := reflect.ValueOf(str)
	reflect.Copy(dst, src)
	fmt.Println("Copy string:", dst.Interface())

	int64Value := int64(9999)
	dst = reflect.New(reflect.ValueOf(int64Value).Type()).Elem()
	src = reflect.ValueOf(int64Value)
	reflect.Copy(dst, src)
	fmt.Println("Copy int64:", dst.Interface())

	float64Value := float64(3.14)
	dst = reflect.New(reflect.ValueOf(float64Value).Type()).Elem()
	src = reflect.ValueOf(float64Value)
	reflect.Copy(dst, src)
	fmt.Println("Copy float64:", dst.Interface())

	boolValue := true
	dst = reflect.New(reflect.ValueOf(boolValue).Type()).Elem()
	src = reflect.ValueOf(boolValue)
	reflect.Copy(dst, src)
	fmt.Println("Copy bool:", dst.Interface())

	sliceValue := []int{1, 2, 3}
	dst = reflect.New(reflect.ValueOf(sliceValue).Type()).Elem()
	src = reflect.ValueOf(sliceValue)
	reflect.Copy(dst, src)
	fmt.Println("Copy slice:", dst.Interface())

	struct1 := struct {
		ID   int
		Name string
	}{ID: 100, Name: "Test Name"}
	dst = reflect.New(reflect.ValueOf(struct1).Type()).Elem()
	src = reflect.ValueOf(struct1)
	reflect.Copy(dst, src)
	fmt.Println("Copy struct:", dst.Interface())
}
