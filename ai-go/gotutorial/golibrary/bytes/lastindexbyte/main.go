package main

import (
	"bytes"
	"fmt"
)

/*
LastIndexByte返回s中c的最后一个实例的索引，如果s中不存在c，则返回-1。
*/
func main() {
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('g'))) // 3
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('r'))) // 8
	fmt.Println(bytes.LastIndexByte([]byte("go gopher"), byte('z'))) // -1
}
