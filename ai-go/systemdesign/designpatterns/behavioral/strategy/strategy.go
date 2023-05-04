package main

import "fmt"

/*
# 策略模式

	定义一系列算法，让这些算法在运行时可以互换，使得分离算法，符合开闭原则。

解释：

	策略模式是指有一定行动内容的相对稳定的策略名称。策略模式在古代中又称“计策”，简称“计”，如《汉书·高帝纪上》：“汉

王从其计”。这里的“计”指的就是计谋、策略。策略模式具有相对稳定的形式，如“避实就虚”、“出奇制胜”等。一定的策略模式，
既可应用于战略决策，也可应用于战术决策；既可实施于大系统的全局性行动，也可实施于大系统的局部性行动。

简介：

	策略模式作为一种软件设计模式，指对象有某个行为，但是在不同的场景中，该行为有不同的实现算法。比如每个人都要“交个人所

得税”，但是“在美国交个人所得税”和“在中国交个人所得税”就有不同的算税方法。

策略模式：

	定义了一组算法（业务规则）；
	封装了每个算法；
	这族的算法可互换代替（interchangeable）。

概念

	定义一组算法，将每个算法都封装起来，并且使它们之间可以互换。

组成：

	抽象策略角色： 策略类，通常由一个接口或者抽象类实现。
	具体策略角色：包装了相关的算法和行为。
	环境角色：持有一个策略类的引用，最终给客户端调用。

使用场景

	多个类只有在算法或行为上稍有不同的场景。
	算法需要自由切换的场景。
	需要屏蔽算法规则的场景。
	具体策略数量超过 4 个，则需要考虑使用混合模式

结构

	Context 封装角色：它也叫做上下文角色，起承上启下封装作用，屏蔽高层模块对策略、算法的直接访问，封装可能存在的变化。
	Strategy 抽象策略角色：策略、算法家族的抽象，通常为接口，定义每个策略或算法必须具有的方法和属性。
	ConcreteStrategy 具体策略角色：实现抽象策略中的操作，该类含有具体的算法。

应用场景：

	1、 多个类只区别在表现行为不同，可以使用Strategy模式，在运行时动态选择具体要执行的行为。
	2、 需要在不同情况下使用不同的策略(算法)，或者策略还可能在未来用其它方式来实现。
	3、 对客户隐藏具体策略(算法)的实现细节，彼此完全独立。
*/
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash\n", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}

func main() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()

	payment = NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
}
