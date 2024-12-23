/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 14:36:55
*/
package main

import (
	"fmt"
)

/*
下面的例子展示了闭包的概念。闭包一般与匿名函数相关，匿名函数可以引用外部函数中定义的变量，对其形成闭包。该变量可作为“半全局变量”，生命周期存在于多次匿名
函数调用中，任何对它的修改都可在匿名函数中可见。
*/
func intSeq() func() int {
	i := 0
	// 匿名函数对外部定义的i形成闭包
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	// 每次调用匿名函数都会将i加1
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
