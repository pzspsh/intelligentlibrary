package main

import "fmt"

/*
# 状态模式

	状态模式用于分离状态和行为。

解释：

	允许一个对象在其内部状态改变时改变它的行为。对象看起来似乎修改了它的类

定义：

	(源于Design Pattern)：当一个对象的内在状态改变时允许改变其行为，这个对象看起来像是改变了其类。
	状态模式主要解决的是当控制一个对象状态的条件表达式过于复杂时的情况。把状态的判断逻辑转移到表示不同状态的一系列类中，

可以把复杂的判断逻辑简化。

概念

	当一个对象内在状态改变时允许其改变行为，这个对象看起来像改变了其类。

意图：

	允许一个对象在其内部状态改变时改变它的行为

使用场景

	行为随状态改变而改变的场景，这也是状态模式的根本出发点，例如权限设计，人员的状态不同即使执行相同的行为结果也会不同，

在这种情况下需要考虑使用状态模式。条件、分支判断语句的替代者

使用场景：

	1.一个对象的行为取决于它的状态，并且它必须在运行时刻根据状态改变它的行为。
	2.一个操作中含有庞大的多分支结构，并且这些分支决定于对象的状态。

结构

	State——抽象状态角色：接口或抽象类，负责对象状态定义，并且封装环境角色以实现状态切换。
	ConcreteState——具体状态角色：每一个具体状态必须完成两个职责：本状态的行为管理以及趋向状态处理，通俗地说，就是本状

态下要做的事情，以及本状态如何过渡到其他状态。

	Context——环境角色：定义客户端需要的接口，并且负责具体状态的切换。
*/
type Week interface {
	Today()
	Next(*DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext {
	return &DayContext{
		today: &Sunday{},
	}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

type Sunday struct{}

func (*Sunday) Today() {
	fmt.Printf("Sunday\n")
}

func (*Sunday) Next(ctx *DayContext) {
	ctx.today = &Monday{}
}

type Monday struct{}

func (*Monday) Today() {
	fmt.Printf("Monday\n")
}

func (*Monday) Next(ctx *DayContext) {
	ctx.today = &Tuesday{}
}

type Tuesday struct{}

func (*Tuesday) Today() {
	fmt.Printf("Tuesday\n")
}

func (*Tuesday) Next(ctx *DayContext) {
	ctx.today = &Wednesday{}
}

type Wednesday struct{}

func (*Wednesday) Today() {
	fmt.Printf("Wednesday\n")
}

func (*Wednesday) Next(ctx *DayContext) {
	ctx.today = &Thursday{}
}

type Thursday struct{}

func (*Thursday) Today() {
	fmt.Printf("Thursday\n")
}

func (*Thursday) Next(ctx *DayContext) {
	ctx.today = &Friday{}
}

type Friday struct{}

func (*Friday) Today() {
	fmt.Printf("Friday\n")
}

func (*Friday) Next(ctx *DayContext) {
	ctx.today = &Saturday{}
}

type Saturday struct{}

func (*Saturday) Today() {
	fmt.Printf("Saturday\n")
}

func (*Saturday) Next(ctx *DayContext) {
	ctx.today = &Sunday{}
}

func main() {
	ctx := NewDayContext()
	todayAndNext := func() {
		ctx.Today()
		ctx.Next()
	}

	for i := 0; i < 8; i++ {
		todayAndNext()
	}
}
