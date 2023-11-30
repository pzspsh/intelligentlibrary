/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:29:25
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建一个 Person 对象
	p1 := Person{Name: "Alice", Age: 30}

	// 将 Person 对象编码为二进制格式
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(p1)

	// 将二进制数据传输给另一个应用程序

	// 在另一个应用程序中，将二进制数据解码为 Person 对象
	decoder := gob.NewDecoder(&buffer)
	var p2 Person
	decoder.Decode(&p2)

	// 打印解码后的 Person 对象
	fmt.Println(p2)
}
