/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:35:33
*/
package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// 编码
	src := []byte("hello")
	maxEnLen := hex.EncodedLen(len(src)) // 最大编码长度
	dst1 := make([]byte, maxEnLen)
	n := hex.Encode(dst1, src)
	dst2 := hex.EncodeToString(src)
	fmt.Println("编码后的结果:", string(dst1[:n]))
	fmt.Println("编码后的结果:", dst2)
	// 解码
	src = dst1
	maxDeLen := hex.DecodedLen(len(src))
	dst1 = make([]byte, maxDeLen)
	n, err := hex.Decode(dst1, src)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s解码后的数据为:%s\n", src, string(dst1[:n]))
	}
	dst3, err := hex.DecodeString(string(src))
	fmt.Printf("%s解码后的数据为:%s\n", src, string(dst3[:n]))
	// dump
	fmt.Printf(hex.Dump(src))
	// dumper
	stdoutDumper := hex.Dumper(os.Stdout)
	defer stdoutDumper.Close()

	stdoutDumper.Write(src)
}

/*
输出内容：
编码后的结果: 68656c6c6f
编码后的结果: 68656c6c6f
68656c6c6f解码后的数据为:hello
68656c6c6f解码后的数据为:hello
00000000  36 38 36 35 36 63 36 63  36 66                    |68656c6c6f|
00000000  36 38 36 35 36 63 36 63  36 66
*/
