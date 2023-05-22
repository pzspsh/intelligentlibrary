/*
@File   : myatoi.go
@Author : pan
@Time   : 2023-05-22 09:41:57
*/
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(myAtoi("42"))
}

// 字符串转换整数 (atoi)
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	flag := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			flag = -1
		} else if v == '+' && i == 0 {
			flag = 1
		} else {
			break
		}
		if result > math.MaxInt32 {
			if flag == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	return flag * result
}
