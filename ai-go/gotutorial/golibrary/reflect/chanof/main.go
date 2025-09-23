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
