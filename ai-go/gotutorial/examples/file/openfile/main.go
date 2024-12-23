/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 10:31:57
*/
package main

import (
	"log"
	"os"
)

func main() {
	// 简单地以只读的方式打开。下面的例子会介绍读写的例子。
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// OpenFile提供更多的选项。
	// 最后一个参数是权限模式permission mode
	// 第二个是打开时的属性
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// 下面的属性可以单独使用，也可以组合使用。
	// 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
	// os.O_CREATE|os.O_APPEND
	// 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
}
