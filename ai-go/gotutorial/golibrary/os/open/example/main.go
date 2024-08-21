/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:12:41
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("path/helloworld/hello.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //必须关闭文件流

	//操作文件
	fmt.Println(file) //&{0xc000100780}
	//读取文件里面的内容
	var strSlice []byte
	var tempSlice = make([]byte, 20)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF { //err == io.EOF表示读取完毕
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		strSlice = append(strSlice, tempSlice[:n]...) //注意写法
	}
	fmt.Println(string(strSlice)) //强制类型转换
}
