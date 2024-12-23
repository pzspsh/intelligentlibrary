/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 09:59:42
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	// 写入缓存
	// byte等同于 int8
	bw.WriteByte('H')
	bw.WriteByte('e')
	bw.WriteByte('l')
	bw.WriteByte('l')
	bw.WriteByte('o')
	bw.WriteByte(' ')
	// rune等同于int32
	bw.WriteRune('世')
	bw.WriteRune('界')
	bw.WriteRune('！')
	// 写入b
	bw.Flush()
	fmt.Println(b)
}
