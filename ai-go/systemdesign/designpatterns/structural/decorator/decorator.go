package main

import "fmt"

/*
# 装饰模式
	装饰模式使用对象组合的方式动态改变或增加对象行为。
	Go语言借助于匿名组合和非入侵式接口可以很方便实现装饰模式。
	使用匿名组合，在装饰器中不必显式定义转调原对象方法。

解释：
	装饰模式指的是在不必改变原类文件和使用继承的情况下，动态地扩展一个对象的功能。它是通过创建一个包装对象，也就是装饰来
包裹真实的对象。

概念
	动态地给一个对象添加一些额外的职责。就增加功能来说，装饰模式相比生成子类更为灵活。

定义：
	23种设计模式之一，英文叫Decorator Pattern，又叫装饰者模式。装饰模式是在不必改变原类文件和使用继承的情况下，动态地扩
展一个对象的功能。它是通过创建一个包装对象，也就是装饰来包裹真实的对象。

特点：
	（1） 装饰对象和真实对象有相同的接口。这样客户端对象就能以和真实对象相同的方式和装饰对象交互。
	（2） 装饰对象包含一个真实对象的引用（reference）
	（3） 装饰对象接受所有来自客户端的请求。它把这些请求转发给真实的对象。
	（4） 装饰对象可以在转发这些请求以前或以后增加一些附加功能。这样就确保了在运行时，不用修改给定对象的结构就可以在外部
增加附加的功能。在面向对象的设计中，通常是通过继承来实现对给定类的功能扩展。

使用场景
	需要扩展一个类的功能，或给一个类增加附加功能。
	需要动态地给一个对象增加功能，这些功能可以再动态地撤销。
	需要为一批的兄弟类进行改装或加装功能，当然是首选装饰模式。

结构
	Component 抽象构件：Component 是一个接口或者是抽象类，就是定义我们最核心的对象，也就是最原始的对象。在装饰模式中，必
然有一个最基本、最核心、最原始的接口或抽象类充当Component 抽象构件。
	ConcreteComponent 具体构件：ConcreteComponent 是最核心、最原始、最基本的接口或抽象类的实现，你要装饰的就是它。
	Decorator 装饰角色：一般是一个抽象类，做什么用呢？实现接口或者抽象方法，它里面可不一定有抽象的方法呀，在它的属性里必
然有一个 private 变量指向 Component 抽象构件。
	具体装饰角色：两个具体的装饰类，你要把你最核心的、最原始的、最基本的东西装饰成其他东西。

适用性：
以下情况使用Decorator模式
	1. 需要扩展一个类的功能，或给一个类添加附加职责。
	2. 需要动态的给一个对象添加功能，这些功能可以再动态的撤销。
	3. 需要增加由一些基本功能的排列组合而产生的非常大量的功能，从而使继承关系变的不现实。
	4. 当不能采用生成子类的方法进行扩充时。一种情况是，可能有大量独立的扩展，为支持每一种组合将产生大量的子类，使得子类
数目呈爆炸性增长。另一种情况可能是因为类定义被隐藏，或类定义不能用于生成子类。
*/

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

func main() {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}
