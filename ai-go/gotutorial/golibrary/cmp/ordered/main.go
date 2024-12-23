/*
@File   : main.go
@Author : pan
@Time   : 2023-11-29 15:21:26
*/
package main

/*
定义如下：
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

Ordered 是对有序类型的约束，包含任何支持操作符 <、<=、>= 和 > 的类型。
如果Go未来的版本添加了新的有序类型，则将修改此接口以将新的有序类型包含进去。

需要注意的是，浮点类型可能包含NaN（“非数字”）类型的值。在将NaN值与任何其它
值（无论是否为 NaN）进行比较时，例如 == 或 < 之类的操作符结果都是 false
*/
func main() {

}
