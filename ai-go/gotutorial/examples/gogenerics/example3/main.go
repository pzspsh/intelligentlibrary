/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 11:47:42
*/
package main

import (
	"fmt"
)

/*
下面的例子是一个对泛型输出的基本例子。函数可以有一个额外的类型参数列表，它使用方括号，
但看起来像一个普通的参数列表：func F[T any](p T) { ... }，代码中的[T any]即为类型参数，
意思是该函数支持任何T类型，当我们调用printSlice[string]([]string{“Hello”，“World”})时，
会被类型推导为string类型，不过在编译器完全可以实现类型推导时，也可以省略显式类型，
如：printSlice([]string{“Hello”，“World”}) ，这样也将会是对的。
*/

func printSlice[T any](s []T) {
	for _, v := range s {
		fmt.Printf("%v ", v)
	}
	fmt.Print("\n")
}

func main() {
	printSlice[int]([]int{1, 2, 3, 4, 5})
	printSlice[float64]([]float64{1.01, 2.02, 3.03, 4.04, 5.05})
	printSlice([]string{"Hello", "World"})
	printSlice[int64]([]int64{5, 4, 3, 2, 1})
}
