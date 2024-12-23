/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:30:27
*/
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type User struct {
	Name string
}

// 编码过程
func (v User) MarshalBinary() ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintln(&b, v.Name)
	return b.Bytes(), nil
}

// 自定义解码过程
func (v *User) UnmarshalBinary(data []byte) error {
	// A simple encoding: plain text.
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.Name)
	return err
}

func main() {
	var network bytes.Buffer
	// 1.创建编码器
	enc := gob.NewEncoder(&network)
	// 2.向编码器中写入数据
	err := enc.Encode(User{Name: "酷走天涯"})
	if err != nil {
		log.Fatal(err)
	}

	// 3.创建解码器
	dec := gob.NewDecoder(&network)
	var user User
	// 4.解码
	dec.Decode(&user)
	fmt.Println(user)
}
