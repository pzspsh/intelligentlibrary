/*
@File   : main.go
@Author : pan
@Time   : 2024-03-04 15:51:26
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 获取当前正在运行程序的路径
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("当前程序路径：", execPath)

	// 获取指定文件的路径（相对于当前工作目录）
	fileRelativePath := "../readfilepath/file.txt" // 替换为实际的文件路径
	absFilePath, err := filepath.Abs(fileRelativePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("指定文件路径：", absFilePath)
	/*
		文件路径获取是更加执行文件路径为基准，
		所以如果执行文件路径和指定文件路径不在同一个磁盘分区中，
		那么获取的绝对路径就会出错，
		所以一般我们使用相对路径时，
		需要确保它们都在同一个磁盘分区中。

		如果打包成二进制文件，和有需要的配置文件，需要放在同一个文件夹中，
		然后通过相对路径来获取文件路径。否则，在打包后，
		由于执行文件路径和配置文件路径不在同一个磁盘分区中，会导致获取文件路径出错。
		或者启动是指定配置文件进行启动
	*/
	f, _ := os.Open(absFilePath)
	defer f.Close()
	data, _ := io.ReadAll(f)
	fmt.Println(string(data))
}
