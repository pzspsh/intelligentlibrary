/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:16:34
*/
package main

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		// recover必须在defer函数中调用
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("After mayPanic()")
}
