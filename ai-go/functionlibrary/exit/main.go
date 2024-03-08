/*
@File   : main.go
@Author : pan
@Time   : 2024-03-07 10:58:10
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	go func() {
		fmt.Println("hello go func")
		// time.Sleep(10 * time.Second)
		time.Sleep(20 * time.Second)
		fmt.Println("end go func")
		os.Exit(1)
	}()
	var a int
	for {
		Demo(&a)
		// a = Demo1(a) // 传参不是传地址(指针)的话，函数内部修改不会影响外部变量
		fmt.Printf("a number: %d\n", a)
		if a == 5 {
			fmt.Println("end demo")
			return
		}
	}
}

func Demo(a *int) {
	fmt.Println("hello demo")
	time.Sleep(2 * time.Second)
	*a++
}

func Demo1(a int) int {
	fmt.Println("hello demo")
	time.Sleep(2 * time.Second)
	a++
	return a
}
