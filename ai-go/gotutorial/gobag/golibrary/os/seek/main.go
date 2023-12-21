/*
@File   : main.go
@Author : pan
@Time   : 2023-12-21 16:25:28
*/
package main

import (
	"fmt"
	"os"
)

/*
功能：Seek设置下一个读或写操作的偏移量offset，根据whence来解析，whence可以理解为参考标准，
0意味着相对于文件的原始位置，1意味着相对于当前偏移量，2意味着相对于文件结尾。它返回新的偏移量和
错误（如果存在）。whence的定义如下(io包下)
*/
func main() {
	s := make([]byte, 10)
	file, _ := os.Open("tmp.txt")
	defer file.Close()
	file.Seek(-12, 2) //从离最后位置12的地方开始
	n, _ := file.Read(s)
	fmt.Println(string(s[:n]))
}
