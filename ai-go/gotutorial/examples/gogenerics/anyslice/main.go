/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 13:57:49
*/
package main

import (
	"fmt"
)

type Person[T any] interface {
	Say(T) T
}

type Student[T any] struct {
	Name string
}

func (s *Student[T]) Say(t T) T {
	fmt.Println(s.Name, "say:", t)
	return t
}

type Teacher[T any] struct {
	Name string
}

func (s *Teacher[T]) Say(t T) T {
	fmt.Println(s.Name, "say:", t)
	return t
}

func main() {
	var p1 Person[string] = &Student[string]{Name: "Tom"}
	fmt.Println("main:", p1.Say("ccc"))
	var p2 Person[int] = &Teacher[int]{Name: "Ann"}
	fmt.Println("main:", p2.Say(123))
}
