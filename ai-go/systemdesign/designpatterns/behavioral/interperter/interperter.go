package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
解释器模式

	解释器模式定义一套语言文法，并设计该语言解释器，使用户能使用特定文法控制解释器行为。
	解释器模式的意义在于，它分离多种复杂功能的实现，每个功能只需关注自身的解释。
	对于调用者不用关心内部的解释器的工作，只需要用简单的方式组合命令就可以。

定义：

	一种设计模式。定义了一个解释器，来解释给定语言和文法的句子。其实质是把语言中的每个符号定

义成一个（对象）类，从而把每个程序转换成一个具体的对象树。

概念

	给定一门语言，定义它的文法的一种表示，并定义一个解释器，该解释器使用该表示来解释语言中的句子。

使用场景

	重复发生的问题可以使用解释器模式
	一个简单语法需要解释的场景

结构

	AbstractExpression——抽象解释器：具体的解释任务由各个实现类完成，具体的解释器分别由TerminalExpression 和 Non-terminalExpression 完成。
	TerminalExpression——终结符表达式：实现与文法中的元素相关联的解释操作，通常一个解释器模式中只有一个终结符表达式，但

有多个实例，对应不同的终结符。

	NonterminalExpression——非终结符表达式：文法中的每条规则对应于一个非终结表达式，非终结符表达式根据逻辑的复杂程度而

增加，原则上每个文法规则都对应一个非终结符表达式。

	Context——环境角色：一般是用来存放文法中各个终结符所对应的具体值，这些信息需要存放到环境角色中，很多情况下我们使用Map

来充当环境角色就足够了。
*/
type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}

type AddNode struct {
	left, right Node
}

func (n *AddNode) Interpret() int {
	return n.left.Interpret() + n.right.Interpret()
}

type MinNode struct {
	left, right Node
}

func (n *MinNode) Interpret() int {
	return n.left.Interpret() - n.right.Interpret()
}

type Parser struct {
	exp   []string
	index int
	prev  Node
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

func (p *Parser) newAddNode() Node {
	p.index++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newMinNode() Node {
	p.index++
	return &MinNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newValNode() Node {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{
		val: v,
	}
}

func (p *Parser) Result() Node {
	return p.prev
}

func main() {
	p := &Parser{}
	p.Parse("1 + 2 + 3 - 4 + 5 - 6")
	res := p.Result().Interpret()
	expect := 1
	if res == expect {
		fmt.Printf("expect %d got %d", expect, res)
	}
}
