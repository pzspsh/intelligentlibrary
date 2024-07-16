/*
@File   : main.go
@Author : pan
@Time   : 2024-07-16 10:50:21
*/
package main

import (
	"fmt"
)

/*
"高内聚" 是软件设计中一个重要的原则，指的是一个模块或组件的功能应该高度集中，相关性强的代码应该被组织在一起。
高内聚的模块通常具有以下特点：

单一职责：一个模块只做一件事，并且做好它。
功能集中：模块内的所有元素（如函数、变量等）都紧密围绕一个中心目标。
低耦合：虽然这与高内聚不是直接相关的，但通常高内聚的模块与其他模块之间的耦合度较低。
现在，让我们用 Go 语言（Golang）来实现一个高内聚的示例。假设我们要实现一个简单的计算器模块，该模块包含加、减、乘、
除四种基本运算。
*/
// Add 实现了加法运算
func Add(a, b float64) float64 {
	return a + b
}

// Subtract 实现了减法运算
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply 实现了乘法运算
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide 实现了除法运算，注意处理除数为0的情况
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

/*
在这个例子中，highcohesion 包是一个高内聚的模块。所有的函数（Add、Subtract、Multiply、Divide）都紧密围绕“计算器”这一中心
目标，每个函数都只做一件事（即一种基本的数学运算），并且这些函数之间没有相互依赖，除了它们都使用了浮点数类型作为参数和返回值。

此外，这个模块与其他模块（如用户界面模块、数据存储模块等）之间的耦合度应该是较低的，因为每个模块都专注于自己的功能，并通过
明确的接口与其他模块进行交互。这种低耦合的设计使得系统更易于理解、维护和扩展。
*/
