/*
@File   : main.go
@Author : pan
@Time   : 2023-11-27 17:18:27
*/
package main

import (
	"bufio"
	"fmt"
	"strings"
)

/*
ScanRunes是Scanner的拆分函数，它返回每个utf -8编码的符文作为令牌。返回的符文序列相当于作为
字符串在输入上的范围循环，这意味着错误的UTF-8编码转换为U+FFFD = "\xef\xbf\xbd"。由于Scan接口，
这使得客户端无法区分正确编码的替换符文和编码错误。
*/
func main() {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}
}
