/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:44:28
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

/*
使用`encoding/gob`包的一个简单用法，创建一个编码器(encoder),传输一些值，然后使用解码器(decoder)进行接收
*/
func main() {
	// 初始化一个encoder和decoder.t通常encoder和decoder将通过网络连接，并且两者在不同的进程中运行
	var network bytes.Buffer        // 使用buffer模拟一个网络连接(二进制字节流)
	enc := gob.NewEncoder(&network) // 编码一些数据到网络中
	dec := gob.NewDecoder(&network) // 从网络中读取编码并解析

	/*
	   //NewDecoder初始化一个decoder对象，返回空的Decoder结构体
	   func NewDecoder(r io.Reader) *Decoder
	   // Decoder结构体方法
	   func (dec *Decoder) Decode(e interface{}) error
	   func (*Decoder) DecodeValue
	   //NewEncoder初始化一个encoder对象，并返回Encoder机构体
	   func NewEncoder(w io.Writer) *Encoder
	   // Encoder结构体方法
	   func (enc *Encoder) Encode(e interface{}) error
	   func (enc *Encoder) EncodeValue(value reflect.Value) error
	*/
	// 使用enc进行发送一些编码的数据
	// 使用enc.Encode方法进行编码两组数据
	err := enc.Encode(P{6, 6, 8, "xxbandy.github.io"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	err = enc.Encode(P{1024, 2048, 1000, "BG彪"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// 使用dec进行解码二进制数据并打印这些值
	// 初始化结构体变量，并将网络连接(&network)中的数据按照结构体Q的实例q进行解码
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 1:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error 2:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}
