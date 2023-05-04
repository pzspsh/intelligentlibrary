package main

import "fmt"

/*
工厂方法模式
	工厂方法模式使用子类的方式延迟生成对象到子类中实现。
	Go中不存在继承 所以使用匿名组合来实现
	Operator 是被封装的实际类接口

解释：
	工厂方法模式（FACTORY METHOD）是一种常用的类创建型设计模式,此模式的核心精神是封装类中变化的部分，提取其中个性化善变
的部分为独立类，通过依赖注入以达到解耦、复用和方便后期维护拓展的目的。它的核心结构有四个角色，分别是抽象工厂；具体工厂；
抽象产品；具体产品

概念
	定义一个用于创建对象的接口，让子类决定实例化哪一个类。工厂方法使一个类的实例化延迟到其子类。

使用场景
	jdbc 连接数据库，硬件访问，降低对象的产生和销毁

模式简介
	工厂方法(Factory Method)模式的意义是定义一个创建产品对象的工厂接口，将实际创建工作推迟到子类当中。核心工厂类不再负责
产品的创建，这样核心类成为一个抽象工厂角色，仅负责具体工厂子类必须实现的接口，这样进一步抽象化的好处是使得工厂方法模式可
以使系统在不修改具体工厂角色的情况下引进新的产品。
	工厂方法模式是简单工厂模式的衍生，解决了许多简单工厂模式的问题。首先完全实现‘开－闭 原则’，实现了可扩展。其次更复
杂的层次结构，可以应用于产品结果复杂的场合。 [2]
	工厂方法模式对简单工厂模式进行了抽象。有一个抽象的Factory类（可以是抽象类和接口），这个类将不再负责具体的产品生产，
而是只制定一些规范，具体的生产工作由其子类去完成。在这个模式中，工厂类和产品类往往可以依次对应。即一个抽象工厂对应一个抽
象产品，一个具体工厂对应一个具体产品，这个具体的工厂就负责生产对应的产品。
	工厂方法模式(Factory Method pattern)是最典型的模板方法模式(Template Method pattern)应用。

角色结构：
	抽象工厂(Creator)角色：是工厂方法模式的核心，与应用程序无关。任何在模式中创建的对象的工厂类必须实现这个接口。
	具体工厂(Concrete Creator)角色：这是实现抽象工厂接口的具体工厂类，包含与应用程序密切相关的逻辑，并且受到应用程序调用
以创建产品对象。在上图中有两个这样的角色：BulbCreator与TubeCreator。
	抽象产品(Product)角色：工厂方法模式所创建的对象的超类型，也就是产品对象的共同父类或共同拥有的接口。在上图中，这个角
色是Light。
	具体产品(Concrete Product)角色：这个角色实现了抽象产品角色所定义的接口。某具体产品有专门的具体工厂创建，它们之间往往
一一对应。
*/

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

// Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

func main() {
	var factory OperatorFactory
	factory = PlusOperatorFactory{}
	if compute(factory, 4, 6) == 10 {
		fmt.Println("模式成功。。。")
	}
	factory = MinusOperatorFactory{}
	if compute(factory, 6, 3) == 3 {
		fmt.Println("模式成功。。。")
	}
}

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
