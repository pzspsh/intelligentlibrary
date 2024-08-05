/*
@File   : main.go
@Author : pan
@Time   : 2024-08-05 16:01:54
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 1、将整数转换成浮点数
	var intvalue int = 18
	floatvalue := float32(intvalue)
	fmt.Println("转换之前:", reflect.TypeOf(intvalue), "将整数转换成浮点数: ", float32(intvalue), "类型：", reflect.TypeOf(floatvalue))

	// 2、将字符串转换成byte
	str := "hello world"
	bytevalue := []byte(str)
	fmt.Println("转换之前:", reflect.TypeOf(str), "将字符串转换成:", bytevalue, "类型：", reflect.TypeOf(bytevalue))

	// 3、将byte转换成string
	bytevalue[0] = 'H'
	str = string(bytevalue)
	fmt.Println("转换之前:", reflect.TypeOf(bytevalue), "将byte转换成string:", str, "类型：", reflect.TypeOf(str))

	// 4、将string转换成rune
	str = "博客"
	runevalue := []rune(str)
	fmt.Println("转换之前:", reflect.TypeOf(str), "将string转换成rune:", runevalue, "类型：", reflect.TypeOf(runevalue))

	// 5、将rune转换成string
	runevalue[0] = '牛'
	str = string(runevalue)
	fmt.Println("转换之前:", reflect.TypeOf(runevalue), "将rune转换成string:", str, "类型：", reflect.TypeOf(str))

	// 6、将byte转换成int
	var b byte = 100
	intvalue = int(b)
	fmt.Println("转换之前:", reflect.TypeOf(b), "将byte转换成int:", intvalue, "类型：", reflect.TypeOf(intvalue))

	/*
		strconv.ParseInt(str,base,bitSize)
		str：要转换的字符串
		base：进位制（2 进制到 36 进制）
		bitSize：指定整数类型（0:int、8:int8、16:int16、32:int32、64:int64）
		返回转换后的结果和转换时遇到的错误
		如果 base 为 0，则根据字符串的前缀判断进位制（0x:16，0:8，其它:10）
	*/
	// 7、将string转换成整数
	str = "103"
	intvalues, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		fmt.Println("转换失败 error: ", err)
	}
	fmt.Println("转换之前:", reflect.TypeOf(str), "将string转换成int:", intvalues, "类型：", reflect.TypeOf(intvalues))
	i, ok := strconv.Atoi("100000")
	if ok == nil {
		fmt.Printf("Atoi, i is %v, type is %v\n", i, reflect.TypeOf(i))
	}

	// 8、将整数转换成字符串
	var i2e int64 = 0x100
	str = strconv.FormatInt(i2e, 10) // FormatInt第二个参数表示进制，10表示十进制。
	fmt.Println(str)
	fmt.Println("转换之前:", reflect.TypeOf(i2e), "将整数转换成字符串:", str, "类型：", reflect.TypeOf(str))

	// 9、AppendInt 将 int 型整数 i 转换为字符串形式，并追加到 []byte 的尾部
	bc := make([]byte, 0)
	bc = strconv.AppendInt(bc, -2048, 16)
	fmt.Println("AppendInt: ", bc)
}
