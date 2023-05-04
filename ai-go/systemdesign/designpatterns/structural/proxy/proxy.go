package main

import "fmt"

/*
代理模式

	代理模式用于延迟处理操作或者在进行实际操作前后进行其它处理。
	代理模式的常见用法有：* 虚代理、* COW代理、* 远程代理、* 保护代理、* Cache 代理、* 防火墙代理、* 同步代理、* 智能指引等等

解释：

	所谓的代理者是指一个类别可以作为其它东西的接口。代理者可以作任何东西的接口：网上连接、存储器中的大对象、文件或其它昂

贵或无法复制的资源。

概念

	为其他对象提供一种代理以控制对这个对象的访问。

简介：

	代理模式的定义：为其他对象提供一种代理以控制对这个对象的访问。在某些情况下，一个对象不适合或者不能直接引用另一个对象

，而代理对象可以在客户端和目标对象之间起到中介的作用。

	著名的代理模式例子为引用计数（英语：reference counting）指针对象。
	当一个复杂对象的多份副本须存在时，代理模式可以结合享元模式以减少存储器用量。典型作法是创建一个复杂对象及多个代理者，

每个代理者会引用到原本的复杂对象。而作用在代理者的运算会转送到原本对象。一旦所有的代理者都不存在时，复杂对象会被移除。

结构

	Subject 抽象主题角色：抽象主题类可以是抽象类也可以是接口，是一个最普通的业务类型定义，无特殊要求。
	RealSubject 具体主题角色：也叫做被委托角色、被代理角色。它才是冤大头，是业务逻辑的具体执行者。
	Proxy 代理主题角色：也叫做委托类、代理类。它负责对真实角色的应用，把所有抽象主题类定义的方法、限制委托给真实主题角色

实现，并且在真实主题角色处理完毕前后做预处理和善后处理工作。

组成：

	抽象角色：通过接口或抽象类声明真实角色实现的业务方法。
	代理角色：实现抽象角色，是真实角色的代理，通过真实角色的业务逻辑方法来实现抽象方法，并可以附加自己的操作。
	真实角色：实现抽象角色，定义真实角色所要实现的业务逻辑，供代理角色调用。

分类

	普通代理：在该模式下，调用者只知代理而不用知道真实的角色是谁，屏蔽了真实角色的变更对高层模块的影响，真实的主题角色想

怎么修改就怎么修改，对高层次的模块没有任何的影响，只要你实现了接口所对应的方法，该模式非常适合对扩展性要求较高的场合。

	强制代理：强制代理的概念就是要从真实角色查找到代理角色，不允许直接访问真实角色。高层模块只要调用 getProxy 就可以访问

真实角色的所有方法，它根本就不需要产生一个代理出来，代理的管理已经由真实角色自己完成。

	区别：普通代理就是我们要知道代理的存在，然后才能访问；强制代理则是调用者直接调用真实角色，而不用关心代理是否存在，其

代理的产生是由真实角色决定的。

	动态代理：根据被代理的接口生成所有的方法，也就是说给定一个接口，动态代理会宣称“我已经实现该接口下的所有方法了”。两

条独立发展的线路。动态代理实现代理的职责，业务逻辑实现相关的逻辑功能，两者之间没有必然的相互耦合的关系。通知从另一个切面
切入，最终在高层模块进行耦合，完成逻辑的封装任务。

	意图：横切面编程，在不改变我们已有代码结构的情况下增强或控制对象的行为。
	首要条件：被代理的类必须要实现一个接口。

结构：

	一个是真正的你要访问的对象(目标类)，一个是代理对象,真正对象与代理
	对象实现同一个接口,先访问代理类再访问真正要访问的对象。
	代理模式分为静态代理、动态代理。
	静态代理是由程序员创建或工具生成代理类的源码，再编译代理类。所谓静态也就是在程序运行前就已经存在代理类的字节码文件，

代理类和委托类的关系在运行前就确定了。

	动态代理是在实现阶段不用关心代理类，而在运行阶段才指定哪一个对象。
*/
type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等。。
	res += "pre:"

	// 调用真实对象
	res += p.real.Do()

	// 调用之后的操作，如缓存结果，对结果进行处理等。。
	res += ":after"

	return res
}

func main() {
	// var sub Subject
	sub := &Proxy{}
	res := sub.Do()
	if res == "pre:real:after" {
		fmt.Println("模式成功")
	}
}
