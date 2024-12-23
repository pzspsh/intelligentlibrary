/*
@File   : main.go
@Author : pan
@Time   : 2023-12-05 17:28:41
*/
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// 打开一个文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 通过反射关闭文件
	value := reflect.ValueOf(file)
	if value.Kind() == reflect.Ptr && value.Elem().Kind() == reflect.Struct {
		if closer, ok := value.Elem().FieldByName("File").Interface().(interface{ Close() error }); ok {
			closer.Close()
			fmt.Println("文件已关闭")
		} else {
			fmt.Println("无法关闭文件")
		}
	}
}
