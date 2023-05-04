package main

import "fmt"

/*
# 访问者模式
	访问者模式可以给一系列对象透明的添加功能，并且把相关代码封装到一个类中。
	对象只要预留访问者接口`Accept`则后期为对象添加功能的时候就不需要改动对象。

解释：
	访问者模式（Visitor Pattern）是GoF提出的23种设计模式中的一种，属于行为模式。据《大话设计模式》中说算是最复杂也是最难
以理解的一种模式了。
	定义（源于GoF《Design Pattern》）：表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素类的前提下定义
作用于这些元素的新操作。
	从定义可以看出结构对象是使用访问者模式必备条件，而且这个结构对象必须存在遍历自身各个对象的方法。这便类似于Java语言当
中的collection概念了。

概念
	封装一些作用于某种数据结构中的各元素的操作，它可以在不改变数据结构的前提下定义作用于这些元素的新的操作。

角色：
	1.Visitor 抽象访问者角色，为该对象结构中具体元素角色声明一个访问操作接口。该操作接口的名字和参数标识了发送访问请求给
具体访问者的具体元素角色，这样访问者就可以通过该元素角色的特定接口直接访问它。
	2.ConcreteVisitor.具体访问者角色，实现Visitor声明的接口。
	3.Element 定义一个接受访问操作(accept())，它以一个访问者(Visitor)作为参数。
	4.ConcreteElement 具体元素，实现了抽象元素(Element)所定义的接受操作接口。
	5.ObjectStructure 结构对象角色，这是使用访问者模式必备的角色。它具备以下特性：能枚举它的元素；可以提供一个高层接口以
允许访问者访问它的元素；如有需要，可以设计成一个复合对象或者一个聚集（如一个列表或无序集合）。

使用场景
	1、一个对象结构包含很多类对象，它们有不同的接口，而你想对这些对象实施一些依赖于其具体类的操作，也就说是用迭代器模式
已经不能胜任的情景。
	2、需要对一个对象结构中的对象进行很多不同并且不相关的操作，而你想避免让这些操作“污染”这些对象的类。
	3、 当该对象结构被很多应用共享时，用Visitor模式让每个应用仅包含需要用到的操作。
	4、定义对象结构的类很少改变，但经常需要在此结构上定义新的操作。改变对象结构类需要重定义对所有访问者的接口，这可能需
要很大的代价。如果对象结构类经常改变，那么可能还是在这些类中定义这些操作较好。

结构
	Visitor—抽象访问者：抽象类或者接口，声明访问者可以访问哪些元素，具体到程序中就是 visit 方法的参数定义哪些对象是可以被访问的。
	ConcreteVisitor—具体访问者：它影响访问者访问到一个类后该怎么干，要做什么事情。
	Element—抽象元素：接口或者抽象类，声明接受哪一类访问者访问，程序上是通过 accept 方法中的参数来定义的。
	ConcreteElement—具体元素：实现 accept 方法，通常是 visitor.visit(this)，基本上都形成了一种模式了。
	ObjectStruture—结构对象：元素产生者，一般容纳在多个不同类、不同接口的容器，如 List、Set、Map 等，在项目中，一般很
少抽象出这个角色。

特点：
	访问者模式把数据结构和作用于结构上的操作解耦合，使得操作集合可相对自由地演化。
	访问者模式适用于数据结构相对稳定算法又易变化的系统。因为访问者模式使得算法操作增加变得容易。若系统数据结构对象易于变
化，经常有新的数据对象增加进来，则不适合使用访问者模式。
	访问者模式的优点是增加操作很容易，因为增加操作意味着增加新的访问者。访问者模式将有关行为集中到一个访问者对象中，其改
变不影响系统数据结构。其缺点就是增加新的数据结构很困难。

*/

type Customer interface {
	Accept(Visitor)
}

type Visitor interface {
	Visit(Customer)
}

type EnterpriseCustomer struct {
	name string
}

type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol) Add(customer Customer) {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol) Accept(visitor Visitor) {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name: name,
	}
}

func (c *EnterpriseCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

type ServiceRequestVisitor struct{}

func (*ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

// only for enterprise
type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}

func main() {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("bob"))
	c.Accept(&ServiceRequestVisitor{})
	fmt.Println("####################################")
	c1 := &CustomerCol{}
	c1.Add(NewEnterpriseCustomer("A company"))
	c1.Add(NewIndividualCustomer("bob"))
	c1.Add(NewEnterpriseCustomer("B company"))
	c1.Accept(&AnalysisVisitor{})
}
