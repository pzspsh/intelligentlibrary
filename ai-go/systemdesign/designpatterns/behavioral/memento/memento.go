package main

import "fmt"

/*
# 备忘录模式

	备忘录模式用于保存程序内部状态到外部，又不希望暴露内部状态的情形。
	程序内部状态使用窄接口传递给外部进行存储，从而不暴露程序实现细节。
	备忘录模式同时可以离线保存内部状态，如保存到数据库，文件等。

解释：

	备忘录模式是一种软件设计模式：在不破坏封闭的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可

将该对象恢复到原先保存的状态。

概念

	在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。

使用场景

	需要保存和恢复数据的相关状态场景。
	提供一个可回滚（rollback）的操作。
	需要监控的副本场景中。
	数据库连接的事务管理就是用的备忘录模式。

结构

	Originator 发起人角色：记录当前时刻的内部状态，负责定义哪些属于备份范围的状态，负责创建和恢复备忘录数据。
	Memento 备忘录角色：负责存储 Originator 发起人对象的内部状态，在需要的时候提供发起人需要的内部状态。
	Caretaker 备忘录管理员角色：对备忘录进行管理、保存和提供备忘录。

备忘录模式（Memento Pattern）又叫做快照模式（Snapshot Pattern）或Token模式，是GoF的23种设计模式之一，属于行为模式。
定义：在不破坏封闭的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可将该对象恢复到原先保存的状态。
涉及角色：

	1.Originator(发起人)：负责创建一个备忘录Memento，用以记录当前时刻自身的内部状态，并可使用备忘录恢复内部状态。Originator

可以根据需要决定Memento存储自己的哪些内部状态。

	2.Memento(备忘录)：负责存储Originator对象的内部状态，并可以防止Originator以外的其他对象访问备忘录。备忘录有两个接口：

Caretaker只能看到备忘录的窄接口，他只能将备忘录传递给其他对象。Originator却可看到备忘录的宽接口，允许它访问返回到先前状
态所需要的所有数据。

	3.Caretaker(管理者):负责备忘录Memento，不能对Memento的内容进行访问或者操作。
*/
type Memento interface{}

type Game struct {
	hp, mp int
}

type gameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int) {
	g.mp += mpDelta
	g.hp += hpDelta
}

func (g *Game) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Load(m Memento) {
	gm := m.(*gameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}

func main() {
	game := &Game{
		hp: 10,
		mp: 10,
	}

	game.Status()
	progress := game.Save()

	game.Play(-2, -3)
	game.Status()

	game.Load(progress)
	game.Status()
}
