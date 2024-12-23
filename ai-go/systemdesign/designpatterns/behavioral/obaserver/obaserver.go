package main

import "fmt"

/*
观察者模式

	观察者模式用于触发联动。
	一个对象的改变会触发其它观察者的相关动作，而此对象无需关心连动对象的具体实现。

解释：

	观察者模式（有时又被称为模型（Model）-视图（View）模式、源-收听者(Listener)模式或从属者模式）是软件设计模式的一种。

在此种模式中，一个目标物件管理所有相依于它的观察者物件，并且在它本身的状态改变时主动发出通知。这通常透过呼叫各观察者所提
供的方法来实现。此种模式通常被用来实现事件处理系统

概念

	定义对象间一种一对多的依赖关系，使得每当一个对象改变状态，则所有依赖于它的对象都会得到通知并被自动更新。

介绍：

	观察者模式是一种对象行为模式。它定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得

到通知并被自动更新。在观察者模式中，主体是通知的发布者，它发出通知时并不需要知道谁是它的观察者，可以有任意数目的观察者订
阅并接收通知。观察者模式不仅被广泛应用于软件界面元素之间的交互，在业务对象之间的交互、权限管理等方面也有广泛的应用。

	观察者模式（Observer）完美的将观察者和被观察的对象分离开。举个例子，用户界面可以作为一个观察者，业务数据是被观察者，

用户界面观察业务数据的变化，发现数据变化后，就显示在界面上。面向对象设计的一个原则是：系统中的每个类将重点放在某一个功能
上，而不是其他方面。一个对象只做一件事情，并且将他做好。观察者模式在模块之间划定了清晰的界限，提高了应用程序的可维护性和重用性。

	观察者设计模式定义了对象间的一种一对多的组合关系，以便一个对象的状态发生变化时，所有依赖于它的对象都得到通知并自动刷新。

使用场景

	关联行为场景。需要注意的是，关联行为是可拆分的，而不是“组合”关系。
	事件多级触发场景。
	跨系统的消息交换场景，如消息队列的处理机制。

使用场景：

	1、当一个抽象模型有两个方面，其中一个方面依赖于另一方面。将这二者封装在独立的对象中以使它们可以各自独立地改变和复用。
	2、当对一个对象的改变需要同时改变其他对象，而不知道具体有多少对象需要被改变。
	3、当一个对象必须通知其他对象，而它又不能假定其他对象是谁。换言之，不希望这些对象是紧密耦合的。

实现方式：

	观察者模式有很多实现方式，从根本上说，该模式必须包含两个角色：观察者和被观察对象。在刚才的例子中，业务数据是被观察对

象，用户界面是观察者。观察者和被观察者之间存在“观察”的逻辑关联，当被观察者发生改变的时候，观察者就会观察到这样的变化，
并且做出相应的响应。如果在用户界面、业务数据之间使用这样的观察过程，可以确保界面和数据之间划清界限，假定应用程序的需求发
生变化，需要修改界面的表现，只需要重新构建一个用户界面，业务数据不需要发生变化。

结构

	Subject 被观察者：定义被观察者必须实现的职责，它必须能够动态地增加、取消观察者。它一般是抽象类或者是实现类，仅仅完成

作为被观察者必须实现的职责：管理观察者并通知观察者。

	Observer 观察者：观察者接收到消息后，即进行 update（更新方法）操作，对接收到的信息进行处理。
	ConcreteSubject 具体的被观察者：定义被观察者自己的业务逻辑，同时定义对哪些事件进行通知。
	ConcreteObserver 具体的观察者：每个观察在接收到消息后的处理反应是不同，各个观察者有自己的处理逻辑。

角色：
1、抽象主题（Subject）：

	它把所有观察者对象的引用保存到一个聚集里，每个主题都可以有任何数量的观察者。抽象主题提供一个接口，可以增加和删除观察者对象。

2、具体主题（Concrete Subject）：

	将有关状态存入具体观察者对象；在具体主题内部状态改变时，给所有登记过的观察者发出通知。

3、抽象观察者（Observer）：

	为所有的具体观察者定义一个接口，在得到主题通知时更新自己。

4、具体观察者（Concrete Observer）：

	实现抽象观察者角色所要求的更新接口，以便使本身的状态与主题状态协调。
*/
type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
}
