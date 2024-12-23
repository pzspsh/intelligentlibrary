package main

import "fmt"

/*
# 组合模式
	组合模式统一对象和对象集，使得使用相同接口使用对象和对象集。
	组合模式常用于树状结构，用于统一叶子节点和树节点的访问，并且可以用于应用某一操作到所有子节点。

解释：
	组合模式，将对象组合成树形结构以表示“部分-整体”的层次结构。

概念：
	将对象组合成树形结构以表示“部分-整体”的层次结构，使得用户对单个对象和组合对象的使用具有一致性。
组合模式概述：
	组合模式(Composite Pattern)
	组合模式使得用户对单个对象和组合对象的使用具有一致性。
	有时候又叫做部分-整体模式，它使我们树型结构的问题中，模糊了简单元素和复杂元素的概念，客户程序可以像处理简单元素一样来
处理复杂元素,从而使得客户程序与复杂元素的内部结构解耦。
	组合模式让你可以优化处理递归或分级数据结构。有许多关于分级数据结构的例子，使得组合模式非常有用武之地。关于分级数据结
构的一个普遍性的例子是你每次使用电脑时所遇到的:文件系统。文件系统由目录和文件组成。每个目录都可以装内容。目录的内容可以
是文件，也可以是目录。按照这种方式，计算机的文件系统就是以递归结构来组织的。如果你想要描述这样的数据结构，那么你可以使用
组合模式Composite。

定义
(GoF《设计模式》)：将对象组合成树形结构以表示“部分整体”的层次结构。组合模式使得用户对单个对象和组合对象的使用具有一致性。
涉及角色：
	1.Component 是组合中的对象声明接口，在适当的情况下，实现所有类共有接口的默认行为。声明一个接口用于访问和管理Component子部件。
	2.Leaf 在组合中表示叶子结点对象，叶子结点没有子结点。
	3.Composite 定义有枝节点行为，用来存储子部件，在Component接口中实现与子部件有关操作，如增加(add)和删除(remove)等。

适用性
以下情况下适用Composite模式：
	1．你想表示对象的部分-整体层次结构
	2．你希望用户忽略组合对象与单个对象的不同，用户将统一地使用组合结构中的所有对象。

使用场景
	维护和展示部分-整体关系的场景，如树形菜单、文件和文件夹管理。
	从一个整体中能够独立出部分模块或功能的场景。
	只要是树形结构，就考虑使用组合模式。

结构
	Component 抽象构件角色：定义参加组合对象的共有方法和属性，可以定义一些默认的行为或属性。
	Leaf 叶子构件：叶子对象，其下再也没有其他的分支，也就是遍历的最小单位。
	Composite 树枝构件：树枝对象，它的作用是组合树枝节点和叶子节点形成一个树形结构

总结：
	组合模式解耦了客户程序与复杂元素内部结构，从而使客户程序可以像处理简单元素一样来处理复杂元素。
	如果你想要创建层次结构，并可以在其中以相同的方式对待所有元素，那么组合模式就是最理想的选择。本章使用了一个文件系统的
例子来举例说明了组合模式的用途。在这个例子中，文件和目录都执行相同的接口，这是组合模式的关键。通过执行相同的接口，你就可
以用相同的方式对待文件和目录，从而实现将文件或者目录储存为目录的子级元素。
*/

type Component interface {
	Parent() Component
	SetParent(Component)
	Name() string
	SetName(string)
	AddChild(Component)
	Print(string)
}

const (
	LeafNode = iota
	CompositeNode
)

func NewComponent(kind int, name string) Component {
	var c Component
	switch kind {
	case LeafNode:
		c = NewLeaf()
	case CompositeNode:
		c = NewComposite()
	}

	c.SetName(name)
	return c
}

type component struct {
	parent Component
	name   string
}

func (c *component) Parent() Component {
	return c.parent
}

func (c *component) SetParent(parent Component) {
	c.parent = parent
}

func (c *component) Name() string {
	return c.name
}

func (c *component) SetName(name string) {
	c.name = name
}

func (c *component) AddChild(Component) {}

func (c *component) Print(string) {}

type Leaf struct {
	component
}

func NewLeaf() *Leaf {
	return &Leaf{}
}

func (c *Leaf) Print(pre string) {
	fmt.Printf("%s-%s\n", pre, c.Name())
}

type Composite struct {
	component
	childs []Component
}

func NewComposite() *Composite {
	return &Composite{
		childs: make([]Component, 0),
	}
}

func (c *Composite) AddChild(child Component) {
	child.SetParent(c)
	c.childs = append(c.childs, child)
}

func (c *Composite) Print(pre string) {
	fmt.Printf("%s+%s\n", pre, c.Name())
	pre += " "
	for _, comp := range c.childs {
		comp.Print(pre)
	}
}

func main() {
	root := NewComponent(CompositeNode, "root")
	c1 := NewComponent(CompositeNode, "c1")
	c2 := NewComponent(CompositeNode, "c2")
	c3 := NewComponent(CompositeNode, "c3")

	l1 := NewComponent(LeafNode, "l1")
	l2 := NewComponent(LeafNode, "l2")
	l3 := NewComponent(LeafNode, "l3")

	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)

	root.Print("")
}
