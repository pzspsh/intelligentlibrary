/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 15:26:35
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 解析64位精度的float
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// 解析64位整数，0代表从string自动推断数字的进制，此处为10进制
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)
	// Atoi是解析10进制整数的简便写法
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
