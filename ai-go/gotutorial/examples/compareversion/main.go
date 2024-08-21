/*
@File   : main.go
@Author : pan
@Time   : 2024-08-20 16:04:40
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// CompareVersions 比较两个版本号的大小
// 如果 v1 > v2，‌返回 1
// 如果 v1 < v2，‌返回 -1
// 如果 v1 == v2，‌返回 0
func CompareVersions(v1, v2 string) int {
	// 使用 "." 分割版本号
	v1 = strings.ReplaceAll(v1, "v", "")
	v2 = strings.ReplaceAll(v2, "v", "")
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	// 获取两个版本号中较长的长度
	length := len(parts1)
	if len(parts2) > length {
		length = len(parts2)
	}

	// 遍历并比较每一部分
	for i := 0; i < length; i++ {
		var num1, num2 int
		var err error
		// 如果第一部分不存在，‌则视为 0
		if i < len(parts1) {
			num1, err = strconv.Atoi(parts1[i])
			if err != nil {
				fmt.Println("版本号格式错误:", err)
				return 0
			}
		}
		// 如果第二部分不存在，‌则视为 0
		if i < len(parts2) {
			num2, err = strconv.Atoi(parts2[i])
			if err != nil {
				fmt.Println("版本号格式错误:", err)
				return 0
			}
		}
		// 比较两部分的大小
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	// 如果所有部分都相等，‌则版本号相等
	return 0
}

func main() {
	fmt.Println(CompareVersions("v2.1.2", "v2.1.1")) // 输出: 1
	fmt.Println(CompareVersions("v1.2", "v1.3"))     // 输出: -1
	fmt.Println(CompareVersions("1.2.3", "1.2.3"))   // 输出: 0
	fmt.Println(CompareVersions("1.2", "1.2.3"))     // 输出: -1
	fmt.Println(CompareVersions("1.2.3", "1.2"))     // 输出: 1
}
