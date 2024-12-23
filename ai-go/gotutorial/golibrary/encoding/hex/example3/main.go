/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:43:38
*/
package main

import (
	"encoding/hex"
	"fmt"
)

func Something() {
	//func EncodeToString(src []byte) string 编码byte字节为16进制字符串
	src := []byte("hello")
	fmt.Println(src)                     //[104 101 108 108 111]
	encodeStr := hex.EncodeToString(src) //68656c6c6f 16进制转换
	fmt.Println(encodeStr)

	//func Encode(dst, src []byte) int
	//func EncodedLen(n int) int
	Welcome := []byte("Gopher!")
	Wdest := make([]byte, hex.EncodedLen(len(Welcome)))
	num := hex.Encode(Wdest, Welcome)
	fmt.Println(Wdest, num) //num=14

	//func DecodeString(s string) ([]byte, error)  解码16进制的字符串为byte类型
	decodeStr, _ := hex.DecodeString(encodeStr)
	fmt.Println(string(decodeStr))

	//func DecodedLen(x int) int  x个byte解码后的长度，一般是x/2
	//func Decode(dst, src []byte) (int, error) 将byte类型的src解码到byte类型的dst中，并且返回dst的长度
	test := []byte("48656c6c6f20476f7068657221")
	dest := make([]byte, hex.DecodedLen(len(test))) //定义一个切片
	num, err := hex.Decode(dest, test)              //转换16进制字符串为byte[]类型，返回切片长度
	if err != nil {
		return
	}
	fmt.Println(num, dest[:num], string(dest), len(dest), cap(dest)) // print 13

	//func Dump(data []byte) string         //返回给定字符串以及字符串相对应的hex dump文件 效果相当于linux命令行下的"hexdump -C filename"
	content := []byte("Go is an open source programming language.")
	fmt.Println(hex.Dump(content))
}

func main() {
	Something()
}
