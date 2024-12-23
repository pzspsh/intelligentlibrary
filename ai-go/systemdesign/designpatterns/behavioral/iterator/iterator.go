package main

import "fmt"

/*
# 迭代器模式
	迭代器模式用于使用相同方式送代不同类型集合或者隐藏集合类型的具体实现。
	可以使用迭代器模式使遍历同时应用送代策略，如请求新对象、过滤、处理对象等。

解释：
	迭代器模式（Iterator），提供一种方法顺序访问一个聚合对象中的各种元素，而又不暴露该对象的内部表示。

简介：
	在面向对象编程里，迭代器模式是一种设计模式，是一种最简单也最常见的设计模式。它可以让用户透过特定的接口巡访容器中的每
一个元素而不用了解底层的实现。此外，也可以实现特定目的版本的迭代器。

概念
	它提供一种方法访问一个容器对象中各个元素，而又不需暴露该对象的内部细节。

意图：
	提供一种方法顺序访问一个聚合对象中各个元素，而又不需暴露该对象的内部表示。

结构
	Iterator 抽象迭代器：抽象迭代器负责定义访问和遍历元素的接口，而且基本上是有固定的 3 个方法：first()获得第一个元素，
next()访问下一个元素，isDone()是否已经访问到底部（Java 叫做 hasNext()方法）。
	ConcreteIterator 具体迭代器：具体迭代器角色要实现迭代器接口，完成容器元素的遍历。
	Aggregate 抽象容器：容器角色负责提供创建具体迭代器角色的接口，必然提供一个类似createIterator()这样的方法，在Java中一
般是 iterator()方法。
	Concrete Aggregate 具体容器：具体容器实现容器接口定义的方法，创建出容纳迭代器的对象。

适用性：
	访问一个聚合对象的内容而无需暴露它的内部表示
	支持对聚合对象的多种遍历
	为遍历不同的聚合结构提供一个统一的接口 [

参与者：
1.Iterator（迭代器）
	迭代器定义访问和遍历元素的接口
2.ConcreteIterator （具体迭代器）
	具体迭代器实现迭代器接口
	对该聚合遍历时跟踪当前位置
3.Aggregate （聚合）
	聚合定义创建相应迭代器对象的接口
4.ConcreteAggregate （具体聚合）
	具体聚合实现创建相应迭代器的接口，该操作返回ConcreteIterator的一个适当的实例

效果：
它支持以不同的方式遍历一个聚合
迭代器简化了聚合的接口
在同一个聚合上可以有多个遍历
*/

type Aggregate interface {
	Iterator() Iterator
}

type Iterator interface {
	First()
	IsDone() bool
	Next() interface{}
}

type Numbers struct {
	start, end int
}

func NewNumbers(start, end int) *Numbers {
	return &Numbers{
		start: start,
		end:   end,
	}
}

func (n *Numbers) Iterator() Iterator {
	return &NumbersIterator{
		numbers: n,
		next:    n.start,
	}
}

type NumbersIterator struct {
	numbers *Numbers
	next    int
}

func (i *NumbersIterator) First() {
	i.next = i.numbers.start
}

func (i *NumbersIterator) IsDone() bool {
	// fmt.Println(i.next, "BBBBBBB", i.numbers.end)
	return i.next > i.numbers.end
}

func (i *NumbersIterator) Next() interface{} {
	if !i.IsDone() {
		next := i.next
		i.next++
		return next
	}
	return nil
}

func IteratorPrint(i Iterator) {
	// !i.IsDone() 的意思是当i.IsDone()为true是结束执行
	for i.First(); !i.IsDone(); {
		c := i.Next()
		fmt.Println(c)
		//fmt.Printf("%#v\n", c)
	}
}

func main() {
	// var aggregate Aggregate
	aggregate := NewNumbers(1, 10)
	IteratorPrint(aggregate.Iterator())
}
