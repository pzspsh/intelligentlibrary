/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 09:52:40
*/
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("Hello 世界！")
	bw := bufio.NewWriter(b)
	bw.ReadFrom(s)
	//bw.Flush()            //ReadFrom无需使用Flush，其自己已经写入．
	fmt.Println(b) // Hello 世界！
}
