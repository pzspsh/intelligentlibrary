/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 16:41:13
*/
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name1" db:"name2"`
}

func main() {
	var s Student
	v := reflect.ValueOf(&s)
	// 类型
	t := v.Type()
	// 获取字段
	f := t.Elem().Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("db"))
}
