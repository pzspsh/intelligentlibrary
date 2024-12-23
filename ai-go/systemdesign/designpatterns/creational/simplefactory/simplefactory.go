package main

import "fmt"

/*
简单工厂模式:
	确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例。
	go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
	NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。
	在这个simplefactory包中只有API 接口和NewAPI函数为包外可见，封装了实现细节。

解释：
	简单工厂模式是属于创建型模式，又叫做静态工厂方法（Static Factory Method）模式，但不属于23种GOF设计模式之一。简单工厂
模式是由一个工厂对象决定创建出哪一种产品类的实例。简单工厂模式是工厂模式家族中最简单实用的模式，可以理解为是不同工厂模式
的一个特殊实现。

实现方式：
	简单工厂模式的实质是由一个工厂类根据传入的参数，动态决定应该创建哪一个产品类（这些产品类继承自一个父类或接口）的实例。
	该模式中包含的角色及其职责
	工厂（Creator）角色
	简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类的创建产品类的方法可以被外界直接调用，创建所需的产品对象。
	抽象产品（Product）角色
	简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
	具体产品（Concrete Product）角色
	是简单工厂模式的创建目标，所有创建的对象都是充当这个角色的某个具体类的实例。
*/

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// HelloAPI is another API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	// api := NewAPI(1)
	// s := api.Say("Tom")
	// if s == "Hi, Tom" {
	// 	fmt.Println("Type1 test correct")
	// }
	api := NewAPI(2)
	s := api.Say("Tom")
	if s == "Hello, Tom" {
		fmt.Println("Type2 test correct")
	}
}
