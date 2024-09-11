/*
@File   : main.go
@Author : pan
@Time   : 2024-09-10 15:09:26
*/
package main

// 判断结构体是否为空
import (
	"fmt"
	"reflect"
)

type A struct {
	name string
	age  int
}

func (a A) IsEmpty() bool {
	return reflect.DeepEqual(a, A{})
}

type MyStruct struct {
	Field1 int
	Field2 string
}

func IsStructZeroValue(s interface{}) bool {
	return reflect.DeepEqual(s, reflect.Zero(reflect.TypeOf(s)).Interface())
}

func RunMain() {
	var a A

	if a == (A{}) { // 括号不能去
		fmt.Println("a == A{} empty")
	}

	if a.IsEmpty() {
		fmt.Println("reflect deep is empty")
	}

	a.name = "pan"
	a.age = 19
	if !a.IsEmpty() {
		fmt.Println("1 reflect deep is not empty")
	} else {
		fmt.Println("1 reflect deep is empty")
	}
}

func RunMain1() {
	// 创建一个结构体实例
	var myStruct MyStruct

	myStruct.Field1 = 42
	// 判断结构体是否为零值
	if IsStructZeroValue(myStruct) {
		fmt.Println("结构体是空的")
	} else {
		fmt.Println("结构体不是空的")
	}
}

func RunMain2() {
	var a A
	a.name = "pan"
	if !reflect.DeepEqual(a, A{}) {
		fmt.Println("struct is not empty")
	} else {
		fmt.Println("struct is empty")
	}
}

func main() {
	RunMain()
	RunMain1()
	RunMain2()
}
