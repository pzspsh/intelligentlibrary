/*
@File   : polymorphism.go
@Author : pan
@Time   : 2023-05-16 15:36:38
*/
package main

import "fmt"

/*
在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口。如果类实现了接口，便可将对象实例赋值给接口。
*/
type Money interface {
	show() string
}

type OldMoney struct {
}

func (oldMoney *OldMoney) show() string {
	return "I am old money"
}

type NewMoney struct {
}

func (newMoney *NewMoney) show() string {
	return "I am new money"
}

func PrintMoney(l []Money) {
	for _, item := range l {
		fmt.Println(item.show())
	}
}

func main() {
	moneyList := []Money{new(OldMoney), new(NewMoney), new(OldMoney)}
	PrintMoney(moneyList)
}
