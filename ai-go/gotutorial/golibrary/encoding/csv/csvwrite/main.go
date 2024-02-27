/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:16:27
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("path/data2.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	writer.Write([]string{"id", "name", "age"})
	writer.WriteAll([][]string{{"1", "Tom", "18"}, {"2", "Jack", "19"}, {"3", "Letty", "20"}})

	if err := writer.Error(); err != nil {
		fmt.Println("执行中发生错误: ", err)
	}
	// 可以调用Flush函数立刻将缓冲区中的数据写入文件中
	writer.Flush()
}
