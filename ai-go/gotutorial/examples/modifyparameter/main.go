/*
@File   : main.go
@Author : pan
@Time   : 2024-04-23 16:30:14
*/
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type person Person
type Perso = Person
type Pers func() Person

func main() {
	pers := Demo()
	pers1 := Demo1()
	pers2 := Demo2()
	fmt.Println(pers)
	fmt.Println(pers1)
	fmt.Println(pers2)
}

func Demo() person {
	per := person(test())
	per.Age = 28
	per.Name = "pan"
	return per
}

func Demo1() Perso {
	per := Pers(test)
	pers := per()
	pers.Age = 28
	pers.Name = "pan1"
	return pers
}

func Demo2() Perso {
	pers := test()
	pers.Age = 28
	pers.Name = "pan2"
	return pers
}

func test() Person {
	person := Person{}
	return person
}
