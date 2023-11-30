/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:34:36
*/
package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func main() {

	srccode := []byte("Hello world!")
	fmt.Printf("srccode(%T) = %v\n", srccode, srccode)

	// 编码部分
	fmt.Println("-------------------------------编码部分---------------------------------")

	srccodeLen := len(srccode)
	fmt.Printf("srccodeLen(%T) = %v\n", srccodeLen, srccodeLen)

	// hex.EncodedLen(len(srcCode)) 返回值实际是 len(srcCode) 长度的2倍
	dstEncode := make([]byte, hex.EncodedLen(len(srccode)))

	// 将 srcCode 编码为十六进制，返回编码后的长度
	encodeLen := hex.Encode(dstEncode, srccode)
	fmt.Printf("encodeLen(%T) = %v\n", encodeLen, encodeLen)
	fmt.Printf("dstEncode(%T) = %v\n", dstEncode, dstEncode)

	// 将 srcCode 编码为字符串
	encodedStr := hex.EncodeToString(srccode)
	fmt.Printf("encodedStr(%T) = %v\n", encodedStr, encodedStr)

	// 解码部分
	fmt.Println("-------------------------------解码部分---------------------------------")

	dstEncodeLen := len(dstEncode)
	fmt.Printf("dstEncodeLen(%T) = %v\n", dstEncodeLen, dstEncodeLen)

	// hex.DecodedLen(len(dstEncode)) 返回值实际是len(dstEncode) 长度一般
	dstDeCode := make([]byte, hex.DecodedLen(len(dstEncode)))

	// 将十六进制解码为字符串，返回解码后的长度
	decodeLen, err := hex.Decode(dstDeCode, dstEncode)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decodeLen(%T) = %v\n", decodeLen, decodeLen)
	fmt.Printf("dstDeCode(%T) = %v\n", dstDeCode, dstDeCode)

	// 返回十六进制字符串表示的字节
	decodedStr, err := hex.DecodeString(string(dstEncode))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("decodedStr(%T) = %v\n", decodedStr, decodedStr)

	fmt.Println("-------------------------------dump部分---------------------------------")

	content := []byte("Go is an open source programming language.")
	// 转储返回一个包含给定数据的十六进制转储的字符串。十六进制转储的格式与hexdump -C命令行上的输出相匹配
	fmt.Printf("dump(%T) =  \n%v\n", hex.Dump(content), hex.Dump(content))

	fmt.Println("-------------------------------dumper部分---------------------------------")
	lines := []string{
		"Go is an open source programming language.",
		"\n",
		"We encourage all Go users to subscribe to golang-announce.",
	}
	//Dumper 返回一个 WriteCloser，它将所有写入数据的十六进制转储写入 w。转储的格式与hexdump -C命令行上的输出相匹配
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	for _, line := range lines {
		stdoutDumper.Write([]byte(line))
	}

}
