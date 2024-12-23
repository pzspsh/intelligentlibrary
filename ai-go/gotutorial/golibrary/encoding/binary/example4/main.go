/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:27:00
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
)

type Message struct {
	Msglen uint32

	Msgid uint32

	Msgdata []byte
}

// 封包函数
func Pack(len uint32, id uint32, data []byte) ([]byte, error) {
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	//获取一个存放bytes的缓冲区，存储字节序列
	dataBuff := bufferPool.Get().(*bytes.Buffer)
	//将数据长度写入字节流
	err := binary.Write(dataBuff, binary.LittleEndian, len)
	checkerr(err)
	//将id写入字节流
	err = binary.Write(dataBuff, binary.LittleEndian, id)
	checkerr(err)
	//将数据内容写入字节流
	err = binary.Write(dataBuff, binary.LittleEndian, data)
	checkerr(err)
	return dataBuff.Bytes(), nil

}

// 解包函数
func Unpack(data []byte) (*Message, error) {
	//这里可以不需要额外创建一个数据缓冲
	//创建一个io。Reader
	boolBuffer := bytes.NewReader(data)
	msg := &Message{}
	//读取数据长度和id
	err := binary.Read(boolBuffer, binary.LittleEndian, &msg.Msglen)
	checkerr(err)
	err = binary.Read(boolBuffer, binary.LittleEndian, &msg.Msgid)
	checkerr(err)
	//数据包限制
	//if
	//
	//}
	return msg, nil
}

func checkerr(err error) {
	if err != nil {
		fmt.Println("数据写入与读取失败")
	}
}

func main() {

}
