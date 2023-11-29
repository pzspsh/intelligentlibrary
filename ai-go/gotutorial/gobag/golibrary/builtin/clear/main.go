/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:12:21
*/
package main

import "fmt"

func main() {
	d := []Data{
		{
			User:   map[int]string{1: "frank", 2: "lucy"},
			Salary: map[string]int{"frank": 1000, "lucy": 2000},
		},
	}
	fmt.Printf("d=%+v\n", d) // d=[{User:map[1:frank 2:lucy] Salary:map[frank:1000 lucy:2000]}]
	clear(d)                 // Go 1.21.0 新增 3 个内置函数详解
	fmt.Printf("d=%+v\n", d) // d=[{User:map[] Salary:map[]}]

	d1 := []Data1{
		{
			User:   "frank",
			Salary: 1000,
		},
	}
	fmt.Printf("d1=%+v\n", d1) // d1=[{User:frank Salary:1000}]
	clear(d1)                  // Go 1.21.0 新增3个内置函数详解 min、max、clear
	fmt.Printf("d1=%+v\n", d1) // d1=[{User: Salary:0}]
}

type Data struct {
	User   map[int]string
	Salary map[string]int
}

type Data1 struct {
	User   string
	Salary int
}
