package main

import "fmt"

/*
抽象工厂模式:

	抽象工厂模式用于生成产品族的工厂，所生成的对象是有关联的。
	如果抽象工厂退化成生成的对象无关联则成为工厂函数模式。
	比如本例子中使用RDB和XML存储订单信息，抽象工厂分别能生成相关的主订单信息和订单详情信息。
	如果业务逻辑中需要替换使用的时候只需要改动工厂函数相关的类就能替换使用不同的存储方式了。
	OrderMainDAO 为订单主记录

解释：

	抽象工厂模式（Abstract Factory Pattern）隶属于设计模式中的创建型模式，用于产品族的构建。抽象工厂是所有形态的工厂模式

中最为抽象和最具一般性的一种形态。抽象工厂是指当有多个抽象角色时使用的一种工厂模式。抽象工厂模式可以向客户端提供一个接口，
使客户端在不必指定产品的具体情况下，创建多个产品族中的产品对象。

	工厂模式中的每一个形态都是针对一定问题的解决方案，工厂方法针对的是多个产品系列结构；而抽象工厂模式针对的是多个产品族

结构，一个产品族内有多个产品系列。

概念

	为创建一组相关或相互依赖的对象提供一个接口，而且无须指定它们的具体类。

使用场景

	一个对象族（或是一组没有任何关系的对象）都有相同的约束。
	涉及不同操作系统的时候，都可以考虑使用抽象工厂模式。

实现原理：

	抽象工厂模式相对于工厂方法模式来说，就是工厂方法模式是针对一个产品系列的，而抽象工厂模式是针对多个产品系列的，即工厂

方法模式是一个产品系列一个工厂类，而抽象工厂模式是多个产品系列一个工厂类。在抽象工厂模式中，客户端不再负责对象的创建，而
是把这个责任丢给了具体的工厂类，客户端只负责对对象的调用，从而明确了各个类的职责。并且当一系列相互关联的产品被设计到一个
工厂类里后，客户端的调用将会变得非常简单，而且，如果要更换这一系列的产品，则只需要更换一个工厂类即可。

	如果客户端需要创建一些产品结构，而这些产品结构又分别属于不同的产品类别，则可以使用抽象工厂模式，抽象工厂模式中抽象工

厂类负责定义创建对象的接口，具体这一系列对象的创建工作由实现抽象工厂的具体工厂类来完成。

角色：

	抽象工厂模式中存在四种角色，分别是抽象工厂角色，具体工厂角色，抽象产品角色，具体产品角色。
	抽象工厂角色：担任这个角色的是工厂方法模式的核心，它是与应用系统商业逻辑无关的。
	具体工厂角色：这个角色直接在客户端的调用下创建产品的实例。这个角色含有选择合适的产品对象的逻辑，而这个逻辑是与应用系

统的商业逻辑紧密相关的。

	抽象产品角色：担任这个角色的类是工厂方法模式所创建的对象的父类，或它们共同拥有的接口。
	具体产品角色：抽象工厂模式所创建的任何产品对象都是某一个具体产品类的实例。这是客户端最终需要的东西，其内部一定充满了

应用系统的商业逻辑。

功能：

	抽象工厂模式的一个主要功能是它能够隔离要生成的具体产品类， 由于这些类的实际类名部被隐藏在工厂内部，因此客户端根本不

需要关心如何对它们进行实例化的细节。每种设计模式都是针对特定问题的解决方案，而抽象工厂模式面临的问题则是当涉及到有多个产
品等级结构寸，如何更好地进行软件体系结构的设计。
*/
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 为订单详情纪录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory DAO 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAP 为关系型数据库的OrderMainDAO实现
type RDBMainDAO struct{}

// SaveOrderMain ...
func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// RDBDetailDAO 为关系型数据库的OrderDetailDAO实现
type RDBDetailDAO struct{}

// SaveOrderDetail ...
func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// RDBDAOFactory 是RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// XMLMainDAO XML存储
type XMLMainDAO struct{}

// SaveOrderMain ...
func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

// XMLDetailDAO XML存储
type XMLDetailDAO struct{}

// SaveOrderDetail ...
func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}

// XMLDAOFactory 是RDB 抽象工厂实现
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

func main() {
	factory := &RDBDAOFactory{}
	getMainAndDetail(factory)
	fmt.Println("####################")
	factory1 := &XMLDAOFactory{}
	getMainAndDetail(factory1)
}

func getMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}
