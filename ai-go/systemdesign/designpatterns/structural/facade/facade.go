package main

import "fmt"

/*
外观模式
	API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。
	facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。

解释：
	外观模式（Facade），亦称“过程模式”。学校课程评价模式之一。美国教育学者斯泰克1967 年在所著《教育评价的外观》中提出。
主张按照描述和判断资料来评价课程，关键的活动是在课程实施的全过程中进行观察和搜集意见，以了解人们对课程的不同看法。这种模
式不限于检查教学的成果，重视描述和判断教学过程中各种复杂、动态的现象和事物。

概念
	要求一个子系统的外部与其内部的通信必须通过一个统一的对象进行。门面模式提供一个高层次的接口，使得子系统更易于使用。

使用场景
	为一个复杂的模块或子系统提供一个供外界访问的接口
	子系统相对独立——外界对子系统的访问只要黑箱操作即可
	预防低水平人员带来的风险扩散

结构
	Facade 门面角色：客户端可以调用这个角色的方法。此角色知晓子系统的所有功能和责任。一般情况下，本角色会将所有从客户端
发来的请求委派到相应的子系统去，也就说该角色没有实际的业务逻辑，只是一个委托类。
	subsystem 子系统角色：可以同时有一个或者多个子系统。每一个子系统都不是一个单独的类，而是一个类的集合。子系统并不知道
门面的存在。对于子系统而言，门面仅仅是另外一个客户端而已。

结构：
	Facade：这个外观类为子系统提供一个共同的对外接口
	Clients：客户对象通过一个外观接口读写子系统中各接口的数据资源。

适用场景：
在以下情况下可以考虑使用外观模式：
	(1)设计初期阶段，应该有意识的将不同层分离，层与层之间建立外观模式。
	(2) 开发阶段，子系统越来越复杂，增加外观模式提供一个简单的调用接口。
	(3) 维护一个大型遗留系统的时候，可能这个系统已经非常难以维护和扩展，但又包含非常重要的功能，为其开发一个外观类，以便
新系统与其交互。
*/

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API is facade interface of facade package
type API interface {
	Test() string
}

// facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

// AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

// BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}

var expect = "A module running\nB module running"

func main() {
	api := NewAPI()
	ret := api.Test()
	if ret == expect {
		fmt.Printf("expect %s, return %s", expect, ret)
	}
}
