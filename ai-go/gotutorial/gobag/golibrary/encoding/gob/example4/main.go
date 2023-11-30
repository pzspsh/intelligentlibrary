/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:45:39
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// Vector 结构体有一些导出的域，这些包不能够被访问，因此我们需要使用`gob`包
// 写一个BinaryMarshal/BinaryUnmarshal方法对来允许我们去发送和接收该类型的数据.
// 这些接口都被定义在了`encoding`包中
// 其实等同于当前包中定义的`GobEncode/GobDecoder`接口

type Vector struct {
	x, y, z int
}

// Vector的MarshalBinary方法
func (v Vector) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text.
	// 一个简单的纯文本编码示例
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// Vector的UnmarshalBinary方法修改了接受者方法，必须接收一个指针
func (v *Vector) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// 使用自定义的encoding和decoding方法来传输数据
func main() {
	// 使用buffer伪造一个网络连接
	var network bytes.Buffer

	// 初始化一个编码器encoder并发送一段数据
	enc := gob.NewEncoder(&network)
	//因为Vector中的元素都是私有变量不能被外部调用,需要默认定义相关的方法
	// 疑问：为啥Vector结构体相关的方法会自动执行内部的MarshalBinary和UnmarshalBinary方法
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatal("encode:", err)
	}

	// 创建一个解码器(decoder)并接收一个值
	dec := gob.NewDecoder(&network)
	var v Vector
	err = dec.Decode(&v)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(v)
}
