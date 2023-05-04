package main

import "fmt"

/*
# 桥接模式
	桥接模式分离抽象部分和实现部分。使得两部分独立扩展。
	桥接模式类似于策略模式，区别在于策略模式封装一系列算法使得算法可以互相替换。
	策略模式使抽象部分和实现部分分离，可以独立变化。

解释：
	桥接模式是将抽象部分与它的实现部分分离，使它们都可以独立地变化。它是一种对象结构型模式，又称为柄体(Handle and Body)
模式或接口(interface)模式。

概念
	将抽象和实现解耦，使得两者可以独立地变化。

使用场景
	不希望或不适用使用继承的场景
	接口或抽象类不稳定的场景
	重用性要求较高的场景

简介：
	在软件系统中，某些类型由于自身的逻辑，它具有两个或多个维度的变化，那么如何应对这种“多维度的变化”？如何利用面向对象
的技术来使得该类型能够轻松的沿着多个方向进行变化，而又不引入额外的复杂度？这就要使用Bridge模式

意图：
	在提出桥梁模式的时候指出，桥梁模式的用意是"将抽象化(Abstraction)与实现化(Implementation)脱耦，使得二者可以独立地变化
"。这句话有三个关键词，也就是抽象化、实现化和脱耦。
抽象化
	存在于多个实体中的共同的概念性联系，就是抽象化。作为一个过程，抽象化就是忽略一些信息，从而把不同的实体当做同样的实体对待。
实现化
	抽象化给出的具体实现，就是实现化。
脱耦
	所谓耦合，就是两个实体的行为的某种强关联。而将它们的强关联去掉，就是耦合的解脱，或称脱耦。在这里，脱耦是指将抽象化和
实现化之间的耦合解脱开，或者说是将它们之间的强关联改换成弱关联。
	将两个角色之间的继承关系改为聚合关系，就是将它们之间的强关联改换成为弱关联。因此，桥梁模式中的所谓脱耦，就是指在一个
软件系统的抽象化和实现化之间使用组合/聚合关系而不是继承关系，从而使两者可以相对独立地变化。这就是桥梁模式的用意。

结构
	Abstraction——抽象化角色：它的主要职责是定义出该角色的行为，同时保存一个对实现化角色的引用，该角色一般是抽象类。
	Implementor——实现化角色：它是接口或者抽象类，定义角色必需的行为和属性。
	RefinedAbstraction——修正抽象化角色：它引用实现化角色对抽象化角色进行修正。
	ConcreteImplementor——具体实现化角色：它实现接口或抽象类定义的方法和属性。

结构：
	可以看出，这个系统含有两个等级结构，也就是：
	由抽象化角色和修正抽象化角色组成的抽象化等级结构。
	由实现化角色和两个具体实现化角色所组成的实现化等级结构。
桥梁模式所涉及的角色有：
	抽象化(Abstraction)角色：抽象化给出的定义，并保存一个对实现化对象的引用。
	修正抽象化(Refined Abstraction)角色：扩展抽象化角色，改变和修正父类对抽象化的定义。
	实现化(Implementor)角色：这个角色给出实现化角色的接口，但不给出具体的实现。必须指出的是，这个接口不一定和抽象化角色
的接口定义相同，实际上，这两个接口可以非常不一样。实现化角色应当只给出底层操作，而抽象化角色应当只给出基于底层操作的更高
一层的操作。
	具体实现化(Concrete Implementor)角色：这个角色给出实现化角色接口的具体实现。
*/

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}

func main() {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	fmt.Println("\n###########################")
	m1 := NewCommonMessage(ViaEmail())
	m1.SendMessage("have a drink?", "bob")
	fmt.Println("\n===========================")
	m2 := NewUrgencyMessage(ViaSMS())
	m2.SendMessage("have a drink?", "bob")
	fmt.Println("\n#############################")
	m3 := NewUrgencyMessage(ViaEmail())
	m3.SendMessage("have a drink?", "bob")
}
