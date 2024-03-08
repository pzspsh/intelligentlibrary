/*
@File   : main.go
@Author : pan
@Time   : 2024-03-08 14:21:27
*/
package main

import (
	"fmt"
)

type Demo struct {
	count int
}

func (d *Demo) Increase() {
	d.count++
}

func (d Demo) Increase1() {
	d.count = 1 // ineffective assignment to field(外勤分配无效)
}

func main() {
	var d Demo
	/*
	   d.Increase()
	   	fmt.Println(d.count) // 1
	*/
	d.Increase1()
	fmt.Println(d.count) // 0
}
