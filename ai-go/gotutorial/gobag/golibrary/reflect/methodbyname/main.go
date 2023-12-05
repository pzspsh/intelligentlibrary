/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:49:28
*/
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) SayHello() {
	fmt.Println("Hello, my name is", u.Name, ",", "I am", u.Age, "years old.")
}

func (u User) SayHi() {
	fmt.Println("Hi, my name is", u.Name, ",", "I am", u.Age, "years old.")
}

func main() {
	u := User{
		Name: "Tom",
		Age:  18,
	}
	fmt.Println(reflect.TypeOf(u))  // 输出: main.User
	fmt.Println(reflect.ValueOf(u)) // 输出: {Tom 18}

	sayHelloMethod := reflect.ValueOf(u).MethodByName("SayHello")
	fmt.Println(sayHelloMethod) // 输出：{0x1b2d7e0}

	sayHelloMethod.Call(nil) // 调用SayHello()方法，输出：Hello, my name is Tom , I am 18 years old.

	// 判断结构体中是否存在指定的方法
	fmt.Println(reflect.ValueOf(u).MethodByName("SayBye").IsValid()) // 输出：false

	// 获取结构体中的所有方法
	methodNum := reflect.TypeOf(u).NumMethod()
	for i := 0; i < methodNum; i++ {
		method := reflect.TypeOf(u).Method(i)
		fmt.Printf("%s: %v\n", method.Name, method.Type)
	}
}
