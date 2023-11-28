/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 10:52:25
*/
package main

import (
	"fmt"
)

func main() {
	a := make(chan int, 10)
	fmt.Println(cap(a)) //10
	fmt.Println(len(a)) //0
	b := make([]int, 2)
	b = append(b, 1)
	fmt.Println(len(b)) //3
	fmt.Println(cap(b)) //4
}

/*
len()是go中使用频率比较高的一个函数，其用来返回Type v的length，其对应的类型以及返回的值如下：

数组：数组长度
数组指针：数组长度
slice/map：slice 或者map中元素个数
string：字符串中字节数
channel:通道中现有数量

func cap(v Type) int

cap()返回的是容器的容量，有时候和len()返回的值是不同的，其对应的类型和返回情况如下：

数组：数组长度
数组指针：数组长度
slice:slice重新分配后能够达到最大长度
channel:分配channel中缓存的大小
通过对比我们可以看到，在面对数组类型以及数组指针时，len和cap的值都是一样的，都是数组长度．
异同点：cap不支持map，string类型．而在slice 和channel，二者获取的值也是不同的，len取得的是现有值，而cap取得的是最大值．例子如下：
*/
