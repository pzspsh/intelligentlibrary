/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:29:15
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(0)
	chType := reflect.ChanOf(reflect.SendDir, t)
	fmt.Println(chType)
	chValue := reflect.New(chType).Elem()
	ch := chValue.Interface().(chan<- int)

	ch <- 1
}
