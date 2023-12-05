/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:44:59
*/
package main

import (
	"fmt"
)

// 定义了基础类型
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// 定义了包装类型
type container struct {
	// 包装类型中包含了基础类型
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	// num变量可以被直接或间接访问
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())
}
