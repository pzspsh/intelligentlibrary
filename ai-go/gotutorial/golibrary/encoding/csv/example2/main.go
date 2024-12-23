/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:20:54
*/
package main

import (
	"encoding/csv"
	"os"
)

func main() {
	f, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("xEFxBBxBF") // 写入UTF-8 BOM，避免使用Microsoft Excel打开乱码
	writer := csv.NewWriter(f)
	writer.Write([]string{"编号", "姓名", "年龄"})
	writer.Write([]string{"1", "张三", "23"})
	writer.Write([]string{"2", "李四", "24"})
	writer.Write([]string{"3", "王五", "25"})
	writer.Write([]string{"4", "赵六", "26"})
	writer.Flush() // 此时才会将缓冲区数据写入
}
