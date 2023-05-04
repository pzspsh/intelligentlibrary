package main

import "fmt"

/*
	Builder 生成器模式：
	       （建造者模式）将一个复杂对象的构建与它表示分离，使得同样的构建过程可以创建不同的表示
	个人想法：建造者的建造流程是在指挥者中，指挥者在用户通知他现在具体的建造者是谁后，
	         建造出对应的产品，建造者中实现了产品的建造细节
	Builder 是生成器接口

解释：

	建造者模式是设计模式的一种，将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

概念

	将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

使用场景

	相同的方法，不同的执行顺序，产生不同的事件结果时，可以采用建造者模式。
	多个部件或零件，都可以装配到一个对象中，但是产生的运行结果又不相同时，则可以使用该模式。
	产品类非常复杂，或者产品类中的调用顺序不同产生了不同的效能，这个时候使用建造者模式非常合适。

结构

	Product 产品类：通常是实现了模板方法模式，也就是有模板方法和基本方法。
	Builder 抽象建造者：规范产品的组建，一般是由子类实现。
	ConcreteBuilder 具体建造者：实现抽象类定义的所有方法，并且返回一个组建好的对象。
	Director 导演类：负责安排已有模块的顺序，然后告诉 Builder 开始建造

角色简介：
实用范围

	1 当创建复杂对象的算法应该独立于该对象的组成部分以及它们的装配方式时。
	2 当构造过程必须允许被构造的对象有不同表示时。

角色
在这样的设计模式中，有以下几个角色：

	1 builder：为创建一个产品对象的各个部件指定抽象接口。
	2 ConcreteBuilder：实现Builder的接口以构造和装配该产品的各个部件，定义并明确它所创建的表示，并提供一个检索产品的接口。
	3 Director：构造一个使用Builder接口的对象。
	4 Product：表示被构造的复杂对象。ConcreteBuilder创建该产品的内部表示并定义它的装配过程，包含定义组成部件的类，包括将

这些部件装配成最终产品的接口。
*/
type Builder interface {
	Part1()
	Part2()
	Part3()
}

type Director struct {
	builder Builder
}

// NewDirector ...
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct Product
func (d *Director) Construct() {
	d.builder.Part1()
	d.builder.Part2()
	d.builder.Part3()
}

type Builder1 struct {
	result string
}

func (b *Builder1) Part1() {
	b.result += "1"
}

func (b *Builder1) Part2() {
	b.result += "2"
}

func (b *Builder1) Part3() {
	b.result += "3"
}

func (b *Builder1) GetResult() string {
	return b.result
}

type Builder2 struct {
	result int
}

func (b *Builder2) Part1() {
	b.result += 1
}

func (b *Builder2) Part2() {
	b.result += 2
}

func (b *Builder2) Part3() {
	b.result += 3
}

func (b *Builder2) GetResult() int {
	return b.result
}

func main() {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res == "123" {
		fmt.Printf("Builder1 expect 123 acture %s\n", res)
	}

	builder2 := &Builder2{}
	director2 := NewDirector(builder2)
	director2.Construct()
	res2 := builder2.GetResult()
	if res2 == 6 {
		fmt.Printf("Builder2 expect 6 acture %d\n", res2)
	}
}
