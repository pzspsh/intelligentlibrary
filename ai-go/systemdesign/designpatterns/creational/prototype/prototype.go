package main

import (
	"fmt"
)

/*
原型模式
	原型模式使对象能复制自身，并且暴露到接口中，使客户端面向接口编程时，不知道接口实际对象的情况下生成新的对象。
	原型模式配合原型管理器使用，使得客户端在不知道具体类的情况下，通过接口管理器得到新的实例，并且包含部分预设定配置。
	Cloneable 是原型对象需要实现的接口

概念
	用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。

使用场景
	资源优化场景：类初始化需要消化非常多的资源，这个资源包括数据、硬件资源等。
	性能和安全要求的场景：通过 new 产生一个对象需要非常繁琐的数据准备或访问权限，则可以使用原型模式。
	一个对象多个修改者的场景：一个对象需要提供给其他对象访问，而且各个调用者可能都需要修改其值时，可以、考虑使用原型模式
拷贝多个对象供调用者使用。

解决问题
	它主要面对的问题是：“某些结构复杂的对象”的创建工作；由于需求的变化，这些对象经常面临着剧烈的变化，但是他们却拥有比
较稳定一致的接口。

优点
	原型模式实际上就是实现 Cloneable 接口，重写 clone（）方法。
	性能优良：原型模式是在内存二进制流的拷贝，要比直接 new 一个对象性能好很多，特别是要在一个循环体内产生大量的对象时，
原型模式可以更好地体现其优点。
	逃避构造函数的约束：这既是它的优点也是缺点，直接在内存中拷贝，构造函数是不会执行的。
*/

type Resume struct {
	name     string
	sex      string
	age      string
	timeArea string
	company  string
}

func (r *Resume) setPersonalInfo(name, sex, age string) {
	if r == nil {
		return
	}
	r.name = name
	r.age = age
	r.sex = sex
}

func (r *Resume) setWorkExperience(timeArea, company string) {
	if r == nil {
		return
	}
	r.company = company
	r.timeArea = timeArea
}

func (r *Resume) display() {
	if r == nil {
		return
	}
	fmt.Println("个人信息：", r.name, r.sex, r.age)
	fmt.Println("工作经历：", r.timeArea, r.company)
}

func (r *Resume) clone() *Resume {
	if r == nil {
		return nil
	}
	new_obj := (*r)
	return &new_obj
}

func NewResume() *Resume {
	return &Resume{}
}

func main() {
	resume := NewResume()

	resume.setPersonalInfo("hclA", "男", "22")
	resume.setWorkExperience("3", "Apple")
	resume.display()

	cloneresume := resume.clone()
	cloneresume.setPersonalInfo("hclB", "女", "22")
	cloneresume.setWorkExperience("3", "HW")
	cloneresume.display()
}
