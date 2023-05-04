package main

import (
	"fmt"
	"strings"
)

/*
中介者模式

	中介者模式封装对象之间互交，使依赖变的简单，并且使复杂互交简单化，封装在中介者中。
	例子中的中介者使用单例模式生成中介者。
	中介者的change使用switch判断类型。

解释：

	一种设计模式。用一个中介对象来封装一系列对象的交互，从而把一批原来可能是交互关系复杂的对象转换成一组松散耦合的中间对

象，以有利于维护和修改。

概念

	用一个中介对象封装一系列的对象交互，中介者使各对象不需要显示地相互作用，从而使其耦合松散，而且可以独立地改变它们之间的交互。

使用场景

	中介者模式适用于多个对象之间紧密耦合的情况，紧密耦合的标准是：在类图中出现了蜘蛛网状结构，即每个类都与其他的类有直接的联系。

结构

	Mediator 抽象中介者角色：抽象中介者角色定义统一的接口，用于各同事角色之间的通信。
	Concrete Mediator 具体中介者角色：具体中介者角色通过协调各同事角色实现协作行为，因此它必须依赖于各个同事角色。
	Colleague 同事角色：每一个同事角色都知道中介者角色，而且与其他的同事角色通信的时候，一定要通过中介者角色协作。每个同

事类的行为分为两种：一种是同事本身的行为，比如改变对象本身的状态，处理自己的行为等，这种行为叫做自发行为（SelfMethod），
与其他的同事类或中介者没有任何的依赖；第二种是必须依赖中介者才能完成的行为，叫做依赖方法（Dep-Method）。
*/
type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image"

	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	GetMediatorInstance().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}

type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}

var mediator *Mediator

func GetMediatorInstance() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}

func main() {
	mediator := GetMediatorInstance()
	mediator.CD = &CDDriver{}
	mediator.CPU = &CPU{}
	mediator.Video = &VideoCard{}
	mediator.Sound = &SoundCard{}

	//Tiggle
	mediator.CD.ReadData()

	if mediator.CD.Data != "music,image" {
		fmt.Printf("CD unexpect data %s", mediator.CD.Data)
	}

	if mediator.CPU.Sound != "music" {
		fmt.Printf("CPU unexpect sound data %s", mediator.CPU.Sound)
	}

	if mediator.CPU.Video != "image" {
		fmt.Printf("CPU unexpect video data %s", mediator.CPU.Video)
	}

	if mediator.Video.Data != "image" {
		fmt.Printf("VidoeCard unexpect data %s", mediator.Video.Data)
	}

	if mediator.Sound.Data != "music" {
		fmt.Printf("SoundCard unexpect data %s", mediator.Sound.Data)
	}
}
