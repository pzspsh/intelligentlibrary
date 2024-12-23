package main

import (
	"fmt"
	"sync"
)

/*
使用懒惰模式的单例模式:

	使用双重检查加锁保证线程安全
	Singleton 是单例模式接口，导出的
	通过该接口可以避免 GetInstance 返回一个包私有类型的指针

解释：

	单例模式，属于创建类型的一种常用的软件设计模式。通过单例模式的方法创建的类在当前进程中只有一个实例（根据需要，也有可

能一个线程中属于单例，如：仅线程上下文内使用同一个实例）

概念

	确保某一个类只有一个实例，而且自行实例化并向整个系统提供这个实例。
	保证一个类仅有一个实例，并提供一个访问它的全局访问点。

简介：

	单例模式是设计模式中最简单的形式之一。这一模式的目的是使得类的一个对象成为系统中的唯一实例。要实现这一点，可以从客户

端对其进行实例化开始。因此需要用一种只允许生成对象类的唯一实例的机制，“阻止”所有想要生成对象的访问。使用工厂方法来限制
实例化过程。这个方法应该是静态方法（类方法），因为让类的实例去生成另一个唯一实例毫无意义

使用场景

	要求生成唯一序列号的环境；
	在整个项目中需要一个共享访问点或共享数据，例如一个Web页面上的计数器，可以不用把每次刷新都记录到数据库中，使用单例模

式保持计数器的值，并确保是线程安全的；

	创建一个对象需要消耗的资源过多，如要访问IO和数据库等资源；
	需要定义大量的静态常量和静态方法（如工具类）的环境，可以采用单例模式（当然，也可以直接声明为static的方式）。
	在 getSingleton 方法前加 synchronized 关键字，也可以在 getSingleton 方法内增加synchronized 来实现。

实现方式：

	通常单例模式在Java语言中，有两种构建方式：
	懒汉式—线程不安全：最基础的实现方式，线程上下文单例，不需要共享给所有线程，也不需要加synchronize之类的锁，以提高性能。
	懒汉式—线程安全：加上synchronize之类保证线程安全的基础上的懒汉模式，相对性能很低，大部分时间并不需要同步
	饿汉式：指全局的单例实例在类装载时构建。
	双检锁式：在懒汉式基础上利用synchronize关键字和volatile关键字确保第一次创建时没有线程间竞争而产生多个实例，仅第一次

创建时同步，性能相对较高

	登记式：作为创建类的全局属性存在，创建类被装载时创建
	枚举式：java中枚举类本身也是一种单例模式

动机：

	对于系统中的某些类来说，只有一个实例很重要，例如，一个系统中可以存在多个打印任务，但是只能有一个正在工作的任务；一个

系统只能有一个窗口管理器或文件系统；一个系统只能有一个计时工具或ID(序号)生成器。如在Windows中就只能打开一个任务管理器。
如果不使用机制对窗口对象进行唯一化，将弹出多个窗口，如果这些窗口显示的内容完全一致，则是重复对象，浪费内存资源；如果这些
窗口显示的内容不一致，则意味着在某一瞬间系统有多个状态，与实际不符，也会给用户带来误解，不知道哪一个才是真实的状态。因此
有时确保系统中某个对象的唯一性即一个类只能有一个实例非常重要。 [2]

	如何保证一个类只有一个实例并且这个实例易于被访问呢？定义一个全局变量可以确保对象随时都可以被访问，但不能防止我们实例

化多个对象。一个更好的解决办法是让类自身负责保存它的唯一实例。这个类可以保证没有其他实例被创建，并且它可以提供一个访问该
实例的方法。这就是单例模式的模式动机。
*/
/*
import (
  "fmt"
  "sync"
  "time"
)

var once sync.Once

func main() {

  //once循环调用firstMethod函数10次,其实只执行一次
  for i := 0; i < 10; i++ {
    once.Do(firstMethod)
    fmt.Println("count:---", i)
  }

  //起10个协程,虽然用once去调secondMethod函数,但该函数不会被执行
  //只打印------i
  for i := 0; i < 10; i++ {
    go func(i int) {
      once.Do(secondMethod)
      fmt.Println("-----", i)
    }(i)
  }
  //主协程等1min,等上面的10个协程执行完
  time.Sleep(1 * time.Second)
}
func firstMethod() {
  fmt.Println("firstMethod")
}
func secondMethod() {
  fmt.Println("secondMethod")
}


懒汉模式
package main

import (
    "fmt"
    "sync"
)

var instanse *singler
var mutex sync.Mutex

type singler struct {
    Name string
}

func NewSingler()*singler {
    if instanse == nil{
        mutex.Lock()
        defer mutex.Unlock()
        if instanse == nil{
            instanse = new(singler)
            instanse.Name = "test"
        }
    }
    return instanse
}

func main() {
    singler := NewSingler()
    fmt.Println(singler.Name)
}


饿汉模式
package mian

import "fmt"

var instanse *singler

type singler struct {
    Name string
}

func NewSingler()*singler{
    return instanse
}

func init() {
    instanse = new(singler)
    instanse.Name = "test"
}


func main() {
    singler := NewSingler()
    fmt.Println(singler.Name)
}
*/
type Singleton interface{ foo() int }

// singleton 是单例模式类，包私有的
type singleton struct {
	a int
}

func (s singleton) foo() int {
	return s.a
}

var (
	instance *singleton
	once     sync.Once
)

// GetInstance 用于获取单例模式对象
func GetInstance() Singleton {
	once.Do(func() { instance = &singleton{a: 13} })
	fmt.Println("实例对象的信息和地址：", instance, &instance)
	return instance
}

const parCount = 100

func main() {
	// ins1 := GetInstance()
	// ins2 := GetInstance()
	// if ins1 == ins2 {
	// 	fmt.Println("instance is equal")
	// }
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			//协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	//关闭channel，所有协程同时开始运行，实现并行(parallel)
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] == instances[i-1] {
			fmt.Println("instance is equal")
		}
	}
}
