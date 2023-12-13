/*
@File   : main.go
@Author : pan
@Time   : 2023-12-13 10:02:28
*/
package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const INT_SIZE = int(unsafe.Sizeof(0)) //64位操作系统，8 bytes

// 判断我们系统中的字节序类型
func systemEdian() {
	var i = 0x01020304
	fmt.Println("&i:", &i)
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))

	if bs[0] == 0x04 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
	fmt.Printf("temp: 0x%x,%v\n", bs[0], &bs[0])
	fmt.Printf("temp: 0x%x,%v\n", bs[1], &bs[1])
	fmt.Printf("temp: 0x%x,%v\n", bs[2], &bs[2])
	fmt.Printf("temp: 0x%x,%v\n", bs[3], &bs[3])

}

func testBigEndian() {

	var testInt int32 = 0x01020304
	fmt.Printf("%d use big endian: \n", testInt)
	testBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)
	fmt.Printf("int32 to bytes: %x \n", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func testLittleEndian() {

	var testInt int32 = 0x01020304
	fmt.Printf("%x use little endian: \n", testInt)
	testBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Printf("int32 to bytes: %x \n", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func main() {
	systemEdian()
	fmt.Println("")
	testBigEndian()
	testLittleEndian()
}

/*
什么是字节序
字节序，又称端序或尾序（英语中用单词：Endianness 表示），在计算机领域中，指电脑内存中或在数字通信链路中，
占用多个字节的数据的字节排列顺序。

在几乎所有的平台上，多字节对象都被存储为连续的字节序列。例如在 Go 语言中，一个类型为int的变量x地址为0x100，
那么其指针&x的值为0x100。且x的四个字节将被存储在内存的0x100, 0x101, 0x102, 0x103位置。

字节的排列方式有两个通用规则:
	大端序（Big-Endian）将数据的低位字节存放在内存的高位地址，高位字节存放在低位地址。这种排列方式与数据用字节表示时的书写顺序一致，
符合人类的阅读习惯。
	小端序（Little-Endian），将一个多位数的低位放在较小的地址处，高位放在较大的地址处，则称小端序。小端序与人类的阅读习惯相反，
	但更符合计算机读取内存的方式，因为CPU读取内存中的数据时，是从低地址向高地址方向进行读取的。

上面的文字描述有点抽象，我们拿一个例子来解释一下字节排列时的大端序和小端序。

在内存中存放整型数值168496141 需要4个字节，这个数值的对应的16进制表示是0X0A0B0C0D，这个数值在用大端序和小端序排列时的在内存中的示
意图如下：


为何要有字节序
	很多人会问，为什么会有字节序，统一用大端序不行吗？答案是，计算机电路先处理低位字节，效率比较高，因为计算都是从低位开始的。所以，
计算机的内部处理都是小端字节序。在计算机内部，小端序被广泛应用于现代 CPU 内部存储数据；而在其他场景，比如网络传输和文件存储则使用
大端序。

Go语言对字节序的处理
	Go 语言存储数据时的字节序依赖所在平台的 CPU，处理大小端序的代码位于 encoding/binary ,包中的全局变量BigEndian用于操作大端序数据，
LittleEndian用于操作小端序数据，这两个变量所对应的数据类型都实现了ByteOrder接口。
*/
