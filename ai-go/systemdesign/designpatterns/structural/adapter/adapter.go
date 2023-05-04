package main

import "fmt"

/*
适配器模式

	适配器模式用于转换一种接口适配另一种接口。
	实际使用中Adapte一般为接口，并且使用工厂函数生成实例。
	在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非入侵式接口特征，其实Adap

ter也适配Adaptee接口。

	Target 是适配的目标接口

解释：

	在计算机编程中，适配器模式（有时候也称包装样式或者包装）将一个类的接口适配成用户所期待的。一个适配允许通常因为接口不

兼容而不能在一起工作的类工作在一起，做法是将类自己的接口包裹在一个已存在的类中。

概念

	将一个类的接口变换成客户端所期待的另一种接口，从而使原本因接口不匹配而无法在一起工作的两个类能够在一起工作。

使用场景

	你有动机修改一个已经投产中的接口时，适配器模式可能是最适合你的模式。比如系统扩展了，需要使用一个已有或新建立的类，但

这个类又不符合系统的接口，怎么办？详细设计阶段不要考虑使用适配器模式，使用主要场景为扩展应用中。

类适配器

	Target 目标角色：该角色定义把其他类转换为何种接口，也就是我们的期望接口。
	Adaptee 源角色：你想把谁转换成目标角色，这个“谁”就是源角色，它是已经存在的、运行良好的类或对象，经过适配器角色的包

装，它会成为一个崭新、靓丽的角色。

	Adapter 适配器角色：适配器模式的核心角色，其他两个角色都是已经存在的角色，而适配器角色是需要新建立的，它的职责非常简

单：把源角色转换为目标角色，怎么转换？通过继承或是类关联的方式。

对象适配器

	不使用多继承或继承的方式，而是使用直接关联，或者称为委托的方式。
	对象适配器和类适配器的区别：
	类适配器是类间继承，对象适配器是对象的合成关系，也可以说是类的关联关系，这是两者的根本区别。实际项目中对象适配器使用

到的场景相对比较多。

解释

	将一个类的接口转换成客户希望的另外一个接口。Adapter模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

基本概念
客户：需要调用我们的代码的对象。

	Adapter模式的宗旨：保留现有类所提供的服务，向客户提供接口，以满足客户的期望。

主要内容
（1）类适配器：

	当客户在接口中定义了他期望的行为时，我们就可以应用适配器模式，提供一个实现该接口的类，并且扩展已有的类，通过创建子类来实现适配。

下面是类适配器的UML图：
（2）对象适配器：

	对象适配器”通过组合除了满足“用户期待接口”还降低了代码间的不良耦合。在工作中推荐使用“对象适配”。

（3） 缺省适配器模式：

	缺省适配器模式是一种特殊的适配器模式，但这个适配器是由一个抽象类实现的，并且在抽象类中要实现目标接口中所规定的所有方

法，但很多方法的实现都是“平庸”的实现，也就是说，这些方法都是空方法。而具体的子类都要继承此抽象类。
*/
type Target interface {
	Request() string
}

// Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// AdapteeImpl 是被适配的目标类
type adapteeImpl struct{}

// SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

// Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

var expect = "adaptee method"

func main() {
	adaptee := NewAdaptee()
	fmt.Println("adaptee", adaptee)
	target := NewAdapter(adaptee)
	res := target.Request()
	fmt.Println("res", res)
	if res != expect {
		fmt.Printf("expect: %s, actual: %s", expect, res)
	}
}
