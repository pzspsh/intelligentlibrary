/*
@File   : main.go
@Author : pan
@Time   : 2024-03-26 10:45:03
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	PanicRecover(myFunc)
	//非内置函数
	fmt.Println("test")
	// 内置函数
	println("test")
	print(1, 2, 3, 4)      //1234
	println(1, 2, 3, 4, 5) //1 2 3 4
}

// append函数用于将一个或多个元素添加到切片(slice)的末尾
// append函数第一个参数slice是一个切片，elems表示切片的元素，其数据类型与切片元素的数据类型必须相同
// append函数往切片添加元素时，如果切片的容量还没满，那么返回的新切片与源切片指向同一个底层数组
// 如果切片的容量已经耗尽，那么append函数会为切片重新分配一个更大容量的底层数组，并将数据复制过去，此时返回的切片则指向新分配的数组
func AppendDemo() {
	s0 := []int{1, 2}

	// 添加一个元素 s1 = []int{1, 2, 3}
	s1 := append(s0, 3)

	// 添加多个元素 s2 = []int{1, 2, 3, 4, 5, 6}
	s2 := append(s1, 4, 5, 6)

	// 将切片添加到另一个切片 s3 = []int{1, 2, 3, 4, 5, 6 , 1 , 2}
	s3 := append(s2, s0...)

	// 截取不同切片再合并 s4 = []int{[4 5 6 3 4 5 6 1 2]}
	s4 := append(s3[3:6], s3[2:]...)

	// 空接口切片
	var t []interface{}
	// t == []interface{}{100, 6.8, "test"}
	t = append(t, 100, 6.8, "test")

	var b []byte
	// 将字符串添加字节数组 b = []byte{'j', 'u', 's','t',”,'g','o'}
	b = append(b, "just go"...)
	fmt.Println("s4:", s4)
	fmt.Println(t)
	fmt.Println(b)

	//创建一个容量为2，长度为0的切片
	s11 := make([]int, 0, 2)

	//向s1添加两个元素
	s22 := append(s11, 3, 4)

	fmt.Printf("s11的地址：%p,容量:%d,长度：%d\n", s11, cap(s11), len(s11)) //s1的地址：0xc0000140a0,容量:2,长度：0
	fmt.Printf("s22的地址：%p,容量:%d,长度：%d\n", s22, cap(s22), len(s22)) //s2的地址：0xc0000140a0,容量:2,长度：2
}

// new函数用于为指定的类型分配内存，并返回一个对应内存的指针
// 使用new分配的一般是普通数据类型(如:int，float32等)或复合类型(比如array，struct)
func NewDemo() {
	i := new(int)
	*i = 10
	fmt.Println(*i) //10

	type User struct {
		ID   int
		Name string
	}

	// 下面使用new创建结构体的语句等同于 u := &User{}
	u := new(User)
	fmt.Println(u)
}

// make函数用于创建并初始化slice、channel和map这样的引用类型
// 与new函数相似，make的第一个参数是所要定义的数据类型，而与new函数不同的是，make函数并不是返回一个指针类型，这是因为slice，channel和map是引用类型
// 当创建切片时，必须传入第二个参数来指定切片的长度
// 也可以传入第三个参数来指定义切片的容量，如果不传第三个参数，则容量与长度相等
func MakeDemo() {
	m := make(map[string]string)
	s := make([]int, 2)
	//容量：5，长度：2
	s1 := make([]int, 2, 5)
	// 未指定容量，因此容量与长度都等3
	s2 := make([]string, 3)
	//无缓冲区
	c1 := make(chan int)
	// 缓冲区大小：2
	c := make(chan int, 2)
	fmt.Println(m, s, s1, s2, c1, c)
}

// len函数用于获得指定类型的长度
// len函数只能用于获得数组，数组指针，切片，channel以及字符串这五种类型的长度。如果是数组或者数组指针，len函数返回的是数组元素数量
func LenDemo() {
	a := [5]int{1, 2, 3, 4, 5}
	p := &a
	fmt.Println(a)
	fmt.Printf("数组长度：%d,指针数组长度：%d\n", len(a), len(p))

	//长度为2，容量为10的切片
	s3 := make([]string, 2, 10)
	fmt.Printf("切片的长度：%d", len(s3)) //2

	fmt.Println(len("学习Go语言"))         //14
	fmt.Println(len([]byte("学习Go语言"))) //14
}

// cap函数用于获得指定类型的容量
// 对于数组或者数组指针，容量与长度是相等的，因此cap函数与len函数的返回值是相等 0 <= len(s) <= cap(s)
// 而对于切片与channel类型，则返回其容量，并且容量大于长度，切片与channel的容量与长度的关系为 0 <= len(s) == cap(s)
func CapDemo() {
	a := [5]int{1, 2, 3, 4, 5}
	s := make([]int, 2, 10)
	fmt.Println(a)
	fmt.Printf("数组的长度：%d,数组的容量：%d\n", len(a), cap(a))
	fmt.Printf("切片的长度：%d,切片的容量：%d\n", len(s), cap(s))
}

// 对于channel类型，len函数返回channel当前缓冲区元素的个数
func ChanDemo() {
	//容量为10
	ch := make(chan int, 10)
	var w sync.WaitGroup
	w.Add(1)
	go func(ch chan int) {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
		w.Done()
	}(ch)
	w.Wait()
	fmt.Println(len(ch)) //3
}

// copy函数用于将源切片src复制到目标切片dst中，并返回复制元素的个数
// 源切片与目标切片的数据类型必须相同，且目标切片必须有对应的容量可以存储复制过来的元素
func CopyDemo() {
	var s1 = make([]int, 1)
	var s2 = []int{1, 2, 3}
	i := copy(s1, s2)

	fmt.Printf("复制了:%d个元素\n", i) // 复制了:1个元素
	fmt.Println(s1)              // [1]
}

// delete函数用于根据指定的key删除对应map类型的元素
func DeleteDemo() {
	var bookstore = map[string]float32{
		"Go入门":     100,
		"Go Web编程": 200,
	}

	delete(bookstore, "Go入门")
	//删除不存在的key，并不会报错
	delete(bookstore, "编程思想")

	fmt.Println(bookstore)
}

func myFunc() {
	panic("触发panic")
}

func PanicRecover(f func()) {
	defer func() {
		msg := recover()
		fmt.Println(msg)
	}()
	f()
}

// close函数用于关闭channel，不过close只能关闭双向channel与用于发送数据的channel，向一个只用于接收数据的channel发送数据触发panic
func CloseDemo() {
	// 双向
	var ch1 chan int
	// 用于发送数据
	var ch2 chan<- int
	// 用于接收数据
	var ch3 <-chan int

	close(ch1)
	close(ch2)
	// close(ch3) //报错
	fmt.Println(ch3)
}

func ComplexRealImag() {
	var c1 complex128 = 2 + 10i
	var c2 = complex(2, 10)

	if c1 == c2 {
		fmt.Println("c1与c2相等")
	}

	fmt.Println(real(c1)) // 2
	fmt.Println(imag(c1)) // 10
}
