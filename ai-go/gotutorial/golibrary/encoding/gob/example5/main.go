/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:46:37
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Point struct {
	X, Y int
}

// 定义一个斜边长度
func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

// 定义毕达哥拉斯接口,其中有很多著名的定理，勾股定理就是其一
type Pythagoras interface {
	Hypotenuse() float64
}

// interfaceEncode 函数编码一个接口类型的值到encoder示例中
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	// 需要再调用的时候优先进行注册指定的类型，否则会失败
	// 需要传送一个指针给接口去编码一个接口类型,如果我们直接去传一个p，将会变成具体类型去代替
	err := enc.Encode(&p)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

// interfaceDecode 从字节流中解码下一个接口类型的值并返回它
// 返回一个Pythagoras 接口类型的值
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	// decode将失败除非在链接中具体的类型已被注册(Point{}需要先使用gob.Register()注册才可以进行解码)
	// 一般情况下我们会在主函数调用中去注册
	var p Pythagoras
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal("decode:", err)
	}
	return p
}

// 该示例演示如何去编码一个接口类型的值
// 和正则类型的区别是注册一个具体的类型而不是实现该接口
func main() {
	// 构造一个网络连接
	var network bytes.Buffer

	// 首先必须为encoder何decoder注册一个具体的类型
	// 随后该具体的类型将发送一个生命去实现该接口
	// func Register(value interface{})
	gob.Register(Point{})

	// 创建一个encoder并发送数据
	enc := gob.NewEncoder(&network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
	}

	// 创建一个decoder解码数据并返回
	dec := gob.NewDecoder(&network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Println(result.Hypotenuse())
	}
}
