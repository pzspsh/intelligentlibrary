/*
@File   : main.go
@Author : pan
@Time   : 2023-11-30 15:19:12
*/
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	ReadCsv("filepath/user.csv")
}

func ReadCsv(filePath string) {
	f, err := os.Open(filePath) // 读取文件
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	fmt.Println(f.Name())

	reader := csv.NewReader(f)
	csvData, err := reader.ReadAll() // 读取全部数据
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvData {
		fmt.Println(line)
	}
}
