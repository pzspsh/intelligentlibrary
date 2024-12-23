/*
@File   : main.go
@Author : pan
@Time   : 2023-11-28 11:50:26
*/
package main

import (
	"bytes"
	"fmt"
)

/*
ToTitle将s视为utf -8编码的字节，并返回一个副本，其中包含映射到其标题大小写的所有Unicode字母。
*/
func main() {
	fmt.Printf("%s\n", bytes.ToTitle([]byte("loud noises"))) // LOUD NOISES
	fmt.Printf("%s\n", bytes.ToTitle([]byte("хлеб")))        // ХЛЕБ
	fmt.Println(string(bytes.ToTitle([]byte("Aaa"))))        // AAA
}
